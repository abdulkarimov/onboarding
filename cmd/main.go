package main

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	repositories.NewBot("6044484605:AAHCjCnvgcmReAG1rImUqGc0tQot1eWfs3o")

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
