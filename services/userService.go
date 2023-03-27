package services

import (
	"github.com/abdulkarimov/onboarding/models"
	"github.com/abdulkarimov/onboarding/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(repositories.GetUsers())
}

func AddUser(c *fiber.Ctx) error {
	u := new(models.User)
	c.BodyParser(&u)
	repositories.AddUser(*u)
	return c.JSON(u)
}

func UpdateUser(c *fiber.Ctx) error {
	u := new(models.User)
	c.BodyParser(&u)
	repositories.UpdateUser(u.ID, *u)
	return c.JSON(u)
}

func DeleteUser(c *fiber.Ctx) error {
	return repositories.DeleteUser(c)
}
