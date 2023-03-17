
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/database"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    database.ConnectDb()
    app := fiber.New()
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders:  "*",
    }))

    setupRoutes(app)
    app.Listen(":3000")
}
