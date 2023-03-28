package database

import (
	"github.com/abdulkarimov/onboarding/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	db, err := gorm.Open(sqlite.Open("Users.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Project{}, &models.Stack{}, &models.Form{}, &models.Question{}, &models.Cabinet{}, &models.Contacts{}, &models.Department{}, &models.Hash{}, &models.Position{}, &models.Role{}, &models.Status{}, &models.User{}, &models.Post{})

	DB = Dbinstance{
		Db: db,
	}

}
