package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetDepartment(c *fiber.Ctx) error {

	department := []models.Department{}
	database.DB.Db.Find(&department)

	return c.Status(200).JSON(department)
}

func AddDepartment(c *fiber.Ctx) error {
	department := new(models.Department)
	c.BodyParser(department)
	database.DB.Db.Create(department)

	return c.Status(200).JSON(department)
}

func UpdateDepartment(c *fiber.Ctx) error {

	oldDepartment := models.Department{}
	department := models.Department{}
	c.BodyParser(&department)

	result := database.DB.Db.First(&oldDepartment, c.Params("ID")).Updates(department)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldDepartment)
	}

	return c.SendStatus(204)
}

func DeleteDepartment(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Department{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
