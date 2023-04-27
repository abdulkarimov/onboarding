package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"log"
)

func SearchByFields(key string) []models.User {
	key = "%" + key + "%"
	log.Println("%" + key + "%")
	var users []models.User

	database.DB.Db.Joins("Position").Find(
		&users,
		"Users.Name LIKE ? OR Position.Name LIKE ?", key, key)

	return users
}
