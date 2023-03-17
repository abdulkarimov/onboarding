package main

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/abdulkarimov/onboarding/notifications"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	notifications.SendNotifyTelegram("newUser@gmail.com", "onboarding.com/accept=sajh#325fa")
	notifications.SendNotifyEmail("newUser@gmail.com", "onboarding.com/accept=sajh#325fa")

	setupRoutes(app)

	app.Listen(":3000")
}
