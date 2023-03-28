package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetProject(c *fiber.Ctx) error {

	project := []models.Project{}
	database.DB.Db.Find(&project)

	return c.Status(200).JSON(project)
}

func AddProject(c *fiber.Ctx) error {
	project := new(models.Project)
	c.BodyParser(project)
	database.DB.Db.Create(project)

	return c.Status(200).JSON(project)
}

func UpdateProject(c *fiber.Ctx) error {

	oldProject := models.User{}
	project := models.User{}
	c.BodyParser(&project)

	result := database.DB.Db.First(&oldProject, c.Params("ID")).Updates(project)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldProject)
	}

	return c.SendStatus(204)
}

func DeleteProject(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Project{}, c.Params("ID"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
