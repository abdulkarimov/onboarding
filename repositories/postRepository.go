package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {

	post := []models.Post{}
	database.DB.Db.Find(&post)

	return c.Status(200).JSON(post)
}

func AddPost(c *fiber.Ctx) error {
	post := new(models.Post)
	c.BodyParser(post)
	database.DB.Db.Create(post)

	return c.Status(200).JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {

	oldPost := models.Post{}
	post := models.Post{}
	c.BodyParser(&post)

	result := database.DB.Db.First(&oldPost, c.Params("id")).Updates(post)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(oldPost)
	}

	return c.SendStatus(204)
}

func DeletePost(c *fiber.Ctx) error {
	result := database.DB.Db.Delete(&models.Post{}, c.Params("id"))
	if result.RowsAffected != 0 {
		return c.SendStatus(200)
	}
	return c.SendStatus(204)
}
