package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	user := []models.User{}
	database.DB.Db.Preload("Project").Preload("Stack").Preload("Position").Preload("Department").Preload("Contacts").Preload("Role").Preload("Status").Find(&user)

	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	c.BodyParser(user)
	database.DB.Db.Create(user)

	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {

	oldUser := models.User{}
	user := models.User{}
	c.BodyParser(&user)

	result := database.DB.Db.First(&oldUser, c.Params("id")).Updates(user)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldUser)
	}

	return c.SendStatus(204)
}

func DeleteUser(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.User{}, c.Params("id"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
