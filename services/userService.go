package services


import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/repositories"
)

func GetUsers(c *fiber.Ctx) error {
	return repositories.GetUsers(c)
}

func AddUser(c *fiber.Ctx) error {
	return repositories.AddUser(c)
}

func UpdateUser( c *fiber.Ctx) error {
	return repositories.UpdateUser(c)
}

func DeleteUser( c *fiber.Ctx) error {
	return repositories.DeleteUser(c)
}