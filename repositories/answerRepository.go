package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetAnswers(c *fiber.Ctx) error {

	answer := []models.Answer{}
	database.DB.Db.Find(&answer)

	return c.Status(200).JSON(answer)
}

func AddAnswer(c *fiber.Ctx) error {
	answer := new(models.Answer)
	c.BodyParser(answer)
	database.DB.Db.Create(answer)

	return c.Status(200).JSON(answer)
}

func UpdateAnswer(c *fiber.Ctx) error {

	oldAnswer := models.Answer{}
	answer := models.Answer{}
	c.BodyParser(&answer)

	result := database.DB.Db.First(&oldAnswer, c.Params("ID")).Updates(answer)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldAnswer)
	}

	return c.SendStatus(204)
}

func DeleteAnswer(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Answer{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
