package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetForms(c *fiber.Ctx) error {

	form := []models.Form{}
	database.DB.Db.Preload(clause.Associations).Find(&form)

	return c.Status(200).JSON(form)
}

func AddForm(c *fiber.Ctx) error {
	form := new(models.Form)
	c.BodyParser(form)
	database.DB.Db.Create(form)

	return c.Status(200).JSON(form)
}

func AddQuestionToForm(c *fiber.Ctx) error {
	questionToForm := new(models.FormQuestion)
	c.BodyParser(questionToForm)
	database.DB.Db.Create(questionToForm)

	return c.Status(200).JSON(questionToForm)
}

func AddUserToForm(c *fiber.Ctx) error {
	formUser := new(models.FormUser)
	c.BodyParser(formUser)
	database.DB.Db.Create(formUser)
	return c.Status(200).JSON(formUser)
}

func UpdateForm(c *fiber.Ctx) error {

	oldForm := models.Form{}
	form := models.Form{}
	c.BodyParser(&form)

	result := database.DB.Db.First(&oldForm, c.Params("ID")).Updates(form)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldForm)
	}

	return c.SendStatus(204)
}

func DeleteForm(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Form{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
