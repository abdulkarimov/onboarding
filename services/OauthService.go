package services

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/abdulkarimov/onboarding/pkg/generator"
	notify "github.com/abdulkarimov/onboarding/pkg/telegram"
	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
	"os"
)

const state = "randomstate"

func Login(c *fiber.Ctx) error {
	if gothUser, err := gf.CompleteUserAuth(c); err == nil {
		c.JSON(gothUser.RawData)
	} else {
		gf.BeginAuthHandler(c)
	}
	return nil
}

func Callback(c *fiber.Ctx) error {
	user, err := gf.CompleteUserAuth(c)
	if err != nil {
		return err
	}
	u := &models.User{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: user.AccessToken,
		IsVerified:  false,
		VerifyToken: generator.RandStringBytes(32),
	}
	notify.Notify.SendNotify(u.Email, os.Getenv("URL")+"/auth/user/verify?token="+u.VerifyToken)
	c.JSON(u)
	database.DB.Db.Create(u)
	return nil
}

func Logout(c *fiber.Ctx) error {
	gf.Logout(c)
	c.Redirect("/")
	return nil
}

func Verify(c *fiber.Ctx) error {
	token := c.Query("token")
	var u models.User
	database.DB.Db.Model(models.User{VerifyToken: token}).First(&u)
	u.IsVerified = true

	c.JSON(u)
	return nil
}
