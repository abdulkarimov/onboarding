package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetPosition(c *fiber.Ctx) error {

	position := []models.Position{}
	database.DB.Db.Find(&position)

	return c.Status(200).JSON(position)
}

func AddPosition(c *fiber.Ctx) error {
	position := new(models.Position)
	c.BodyParser(position)
	database.DB.Db.Create(position)

	return c.Status(200).JSON(position)
}

func UpdatePosition(c *fiber.Ctx) error {

	oldPosition := models.Position{}
	position := models.Position{}
	c.BodyParser(&position)

	result := database.DB.Db.First(&oldPosition, c.Params("ID")).Updates(position)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldPosition)
	}

	return c.SendStatus(204)
}

func DeletePosition(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Position{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
