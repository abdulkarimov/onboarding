package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetContacts(c *fiber.Ctx) error {

	contacts := []models.Contacts{}
	database.DB.Db.Find(&contacts)

	return c.Status(200).JSON(contacts)
}

func AddContacts(c *fiber.Ctx) error {
	contacts := new(models.Contacts)
	c.BodyParser(contacts)
	database.DB.Db.Create(contacts)

	return c.Status(200).JSON(contacts)
}

func UpdateContacts(c *fiber.Ctx) error {

	oldContacts := models.Contacts{}
	contacts := models.Contacts{}
	c.BodyParser(&contacts)

	result := database.DB.Db.First(&oldContacts, c.Params("ID")).Updates(contacts)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldContacts)
	}

	return c.SendStatus(204)
}

func DeleteContacts(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Contacts{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
