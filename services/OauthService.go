package services

import (
	"fmt"
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	notify "github.com/abdulkarimov/onboarding/pkg/telegram"
	"github.com/abdulkarimov/onboarding/repositories"
	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
	"os"
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
	u := models.User{
		Contacts: models.Contacts{Email: user.Email},
		Name:     user.Name,
		Img:      user.AvatarURL,
		Info:     user.Description,
	}

	var du = repositories.FindPreloadUser("Contacts", fmt.Sprintf("Contacts.Email = '%s'", user.Email))

	if du.Contacts.Email != user.Email {
		database.DB.Db.Create(&u)
		notify.Notify.SendTelegram(u.Contacts.Email, os.Getenv("URL")+"/auth/user/verify?token="+u.Contacts.Email)
	} else if !du.Verified {
		notify.Notify.SendTelegram(u.Contacts.Email, os.Getenv("URL")+"/auth/user/verify?token="+du.Contacts.Email)
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
	uemail := c.Query("token")
	u := repositories.FindPreloadUser("Contacts", fmt.Sprintf("Contacts.Email = '%s'", uemail))
	repositories.UpdateColumnUser("Verified", "true", u.ID)

	c.JSON(map[int]string{1: uemail, 2: u.Contacts.Email})
	return nil
}