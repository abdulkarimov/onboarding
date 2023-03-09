
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/database"
    "github.com/gofiber/template/html"
)

func main() {
        // Fiber instance
     
    
        // Serve static assets
     

    database.ConnectDb()
    app := fiber.New(fiber.Config{
        Views: html.New("./views", ".html"),
    })
    setupRoutes(app)

    app.Listen(":3000")
}