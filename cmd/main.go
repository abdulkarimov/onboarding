package main

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/pkg/telegram"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	database.ConnectDb()
	notify.New(os.Getenv("TELEGRAM_BOT_TOKEN"))

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
