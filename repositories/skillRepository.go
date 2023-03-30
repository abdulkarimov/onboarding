package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetSkills(c *fiber.Ctx) error {

	skill := []models.Skill{}
	database.DB.Db.Find(&skill)

	return c.Status(200).JSON(skill)
}

func AddSkill(c *fiber.Ctx) error {
	skill := new(models.Skill)
	c.BodyParser(skill)
	database.DB.Db.Create(skill)

	return c.Status(200).JSON(skill)
}

func UpdateSkill(c *fiber.Ctx) error {

	oldSkill := models.Skill{}
	skill := models.Skill{}
	c.BodyParser(&skill)

	result := database.DB.Db.First(&oldSkill, c.Params("ID")).Updates(skill)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldSkill)
	}

	return c.SendStatus(204)
}

func DeleteSkill(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Skill{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
