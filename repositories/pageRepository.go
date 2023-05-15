package repositories

import (
	"math"
	"strconv"

	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetPaginationData(c *fiber.Ctx) error {
	perPage := 10
	page := 1
	pageStr := c.Params("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	offset := (page - 1) * perPage

	var totalRows int64
	database.DB.Db.Preload(clause.Associations).Model(&models.User{}).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(perPage)))

	user := []models.User{}
	PaginationData := map[string]interface{}{
		"nextPage":     page + 1,
		"previousPage": page - 1,
		"currentPage":  page,
		"fourAfter":    page + 4,
		"fourBefore":   page - 4,
		"totalPages":   int(totalPages),
	}
	database.DB.Db.Preload(clause.Associations).Offset(offset).Order("Name").Find(&user)
	return c.Status(200).JSON(fiber.Map{"user": user, "pagination": PaginationData})
}
