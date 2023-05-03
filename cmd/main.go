package main

import (
	"github.com/abdulkarimov/onboarding/database"
	notify "github.com/abdulkarimov/onboarding/pkg/notification"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/sqlite3"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/yandex"
	gf "github.com/shareed2k/goth_fiber"
	"os"
	"time"
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
	gf.SessionStore = session.New(session.Config{
		Expiration:     720 * time.Hour,
		Storage:        sqlite3.New(),
		KeyLookup:      "cookie:_auth_sesion",
		CookieDomain:   "127.0.0.1",
		CookieSecure:   os.Getenv("ENVIRONMENT") == "production",
		CookieHTTPOnly: true, // Should always be enabled
		CookieSameSite: "Lax",
		KeyGenerator:   utils.UUIDv4,
	})

	app := fiber.New()

	setupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	app.Listen(":3000")
}
