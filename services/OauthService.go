package services

import (
	"github.com/abdulkarimov/onboarding/models"
	"github.com/abdulkarimov/onboarding/repositories"
	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
	"log"
)

const state = "randomstate"

var VerifyUser string

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
		VerifyToken: repositories.RandStringBytes(32),
	}
	log.Println(u)
	c.JSON(u)
	return nil
}

func Logout(c *fiber.Ctx) error {
	gf.Logout(c)
	c.Redirect("/")
	return nil
}

func Verify(c *fiber.Ctx) error {
	token := c.Query("vKey")
	if token == VerifyUser {

	}
	return nil
}
