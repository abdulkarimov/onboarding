package main

import (
	"os"

	"github.com/abdulkarimov/onboarding/database"
	notify "github.com/abdulkarimov/onboarding/notification"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/yandex"
)

func main() {
	godotenv.Load(".env")
	database.ConnectDb()
	notify.New()

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_OAUTH_CLIENT_ID"), os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			os.Getenv("GOOGLE_OAUTH_CALLBACK_URL"),
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		),
		yandex.New(os.Getenv("YANDEX_OAUTH_CLIENT_ID"),
			os.Getenv("YANDEX_OAUTH_CLIENT_SECRET"),
			os.Getenv("YANDEX_OAUTH_CALLBACK_URL"),
		),
	)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	setupRoutes(app)
	app.Listen(":3000")
}
