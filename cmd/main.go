package main

import (
	"github.com/abdulkarimov/onboarding/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000",
	}))

	app.Listen(":3000")
}
