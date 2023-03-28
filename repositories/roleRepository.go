package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetRole(c *fiber.Ctx) error {

	role := []models.Role{}
	database.DB.Db.Find(&role)

	return c.Status(200).JSON(role)
}

func AddRole(c *fiber.Ctx) error {
	role := new(models.Role)
	c.BodyParser(role)
	database.DB.Db.Create(role)

	return c.Status(200).JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {

	oldRole := models.Role{}
	role := models.Role{}
	c.BodyParser(&role)

	result := database.DB.Db.First(&oldRole, c.Params("ID")).Updates(role)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldRole)
	}

	return c.SendStatus(204)
}

func DeleteRole(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Role{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
