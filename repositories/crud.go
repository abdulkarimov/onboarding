package repositories

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/models"
	"log"
)

func GetUsers() []models.User {
	var users []models.User
	database.DB.Db.Preload("Contacts").Find(&users)
	return users
}

func FindUser(u models.User) models.User {
	user := models.User{}
	database.DB.Db.Find(&user, u)
	return user
}

func FindPreloadUser(preload string, query string) models.User {
	user := models.User{}
	database.DB.Db.Joins(preload).First(&user, query)
	return user
}

func AddUser(u models.User) {
	database.DB.Db.Create(&u)
}

func RemoveUser(ids []int) {
	database.DB.Db.Delete(&models.User{}, ids)
}
func UpdateColumnUser(cname string, value string, id uint) {
	var u models.User
	database.DB.Db.Find(&u, "id = ?", id).Update(cname, value)
	log.Println(cname, value, id)
}
func UpdateUser(id uint, user models.User) {
	oldUser := models.User{}
	database.DB.Db.First(&oldUser, id).Updates(user)
}
