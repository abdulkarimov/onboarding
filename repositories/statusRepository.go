package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetStatus(c *fiber.Ctx) error {

	status := []models.Status{}
	database.DB.Db.Find(&status)

	return c.Status(200).JSON(status)
}

func AddStatus(c *fiber.Ctx) error {
	status := new(models.Status)
	c.BodyParser(status)
	database.DB.Db.Create(status)

	return c.Status(200).JSON(status)
}

func UpdateStatus(c *fiber.Ctx) error {

	oldStatus := models.Status{}
	status := models.Status{}
	c.BodyParser(&status)

	result := database.DB.Db.First(&oldStatus, c.Params("id")).Updates(status)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldStatus)
	}

	return c.SendStatus(204)
}

func DeleteStatus(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Status{}, c.Params("id"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
