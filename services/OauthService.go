package services

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	notify "github.com/abdulkarimov/onboarding/pkg/telegram"
	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
	"os"
	"strconv"
)

const state = "randomstate"

func Login(c *fiber.Ctx) error {
	return gf.BeginAuthHandler(c)
}

func Callback(c *fiber.Ctx) error {
	user, err := gf.CompleteUserAuth(c)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(user.UserID)
	u := models.User{
		Contacts: models.Contacts{Email: user.Email},
		Name:     user.Name,
		Img:      user.AvatarURL,
		Info:     user.Description,
		HashID:   uint(userId),
	}

	var du models.User
	database.DB.Db.Find(&du, &models.User{HashID: uint(userId)})

	if du.HashID != u.HashID {
		database.DB.Db.Create(&u)
		notify.Notify.SendNotify(u.Contacts.Email, os.Getenv("URL")+"/auth/user/verify?token="+strconv.Itoa(int(u.HashID)))
	} else if du.StatusID == 1 {
		notify.Notify.SendNotify(u.Contacts.Email, os.Getenv("URL")+"/auth/user/verify?token="+strconv.Itoa(int(u.HashID)))
	}

	c.JSON(du)
	return nil
}

func Logout(c *fiber.Ctx) error {
	err := gf.Logout(c)
	if err != nil {
		return err
	}
	err = c.Redirect("/")
	if err != nil {
		return err
	}
	return nil
}

func Verify(c *fiber.Ctx) error {
	token, err := strconv.Atoi(c.Query("token"))
	if err != nil {
		return err
	}
	var u models.User
	database.DB.Db.Find(&u, &models.User{HashID: uint(token)}).Update("StatusID", 1)
	c.JSON(map[int]uint{1: uint(token), 2: u.HashID})
	return nil
}
