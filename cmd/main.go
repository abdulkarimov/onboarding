package main

import (
	"github.com/abdulkarimov/onboarding/database"
	notify "github.com/abdulkarimov/onboarding/pkg/notification"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/yandex"
	"os"
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

	setupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:3000",
		AllowCredentials: true,
	}))

	app.Listen(":3000")
}
