package services

import (
	"github.com/abdulkarimov/onboarding/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return repositories.GetUsers(c)
}

func AddUser(c *fiber.Ctx) error {
	return repositories.AddUser(c)
}

func UpdateUser(c *fiber.Ctx) error {
	return repositories.UpdateUser(c)
}

func DeleteUser(c *fiber.Ctx) error {
	return repositories.DeleteUser(c)
}

func SearchUsers(c *fiber.Ctx) error {
	c.JSON(repositories.SearchByFields(c.Query("key")))
	return nil
}
