package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/getUsers", handlers.GetUsers)
    app.Post("/addUser", handlers.AddUser)
}