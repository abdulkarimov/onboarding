package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
)

func SearchByFields(key string) []models.User {
	key = "%" + key + "%"
	var users []models.User
	var projects []models.Project
	var stacks []models.Stack
	database.DB.Db.Joins("Position").Where("users.name LIKE ?", key).Or(
		"Position.Name LIKE ?", key).Find(&users)
	database.DB.Db.Table("Projects").Preload("Users").Where("Projects.Name LIKE ?", key).Find(&projects)
	database.DB.Db.Table("Stacks").Preload("Users").Where("Stacks.Name LIKE ?", key).Find(&stacks)
	for _, project := range projects {
		users = append(users, project.Users...)
	}
	for _, stack := range stacks {
		users = append(users, stack.Users...)
	}
	return users
}
