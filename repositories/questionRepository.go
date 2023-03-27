package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetQuestion(c *fiber.Ctx) error {

	question := []models.Question{}
	database.DB.Db.Find(&question)

	return c.Status(200).JSON(question)
}

func AddQuestion(c *fiber.Ctx) error {
	question := new(models.Question)
	c.BodyParser(question)
	database.DB.Db.Create(question)

	return c.Status(200).JSON(question)
}

func UpdateQuestion(c *fiber.Ctx) error {

	oldQuestion := models.Question{}
	question := models.Question{}
	c.BodyParser(&question)

	result := database.DB.Db.First(&oldQuestion, c.Params("id")).Updates(question)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldQuestion)
	}

	return c.SendStatus(204)
}

func DeleteQuestion(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Question{}, c.Params("id"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
