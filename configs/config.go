package configs

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/sqlite3"
	"os"
	"time"
)

func GoogleConfig() session.Config {
	config := session.Config{
		Expiration:     30 * time.Minute,
		Storage:        sqlite3.New(), // From github.com/gofiber/storage/sqlite3
		KeyLookup:      "header:session_id",
		CookieDomain:   "google.com",
		CookiePath:     "/users",
		CookieSecure:   os.Getenv("ENVIRONMENT") == "production",
		CookieHTTPOnly: true, // Should always be enabled
		CookieSameSite: "Lax",
		KeyGenerator:   utils.UUIDv4,
	}
	return config
}
