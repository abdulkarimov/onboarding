package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetStacks(c *fiber.Ctx) error {

	stack := []models.Stack{}
	database.DB.Db.Find(&stack)

	return c.Status(200).JSON(stack)
}

func AddStack(c *fiber.Ctx) error {
	stack := new(models.Stack)
	c.BodyParser(stack)
	database.DB.Db.Create(stack)

	return c.Status(200).JSON(stack)
}

func UpdateStack(c *fiber.Ctx) error {

	oldStack := models.Stack{}
	stack := models.Stack{}
	c.BodyParser(&stack)

	result := database.DB.Db.First(&oldStack, c.Params("ID")).Updates(stack)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldStack)
	}

	return c.SendStatus(204)
}

func DeleteStack(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Stack{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
