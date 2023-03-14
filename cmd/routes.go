package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/services"
    "github.com/abdulkarimov/onboarding/validations"
)

func setupRoutes(app *fiber.App) {
    api := app.Group("/api")  

    //user
    api.Get("/getUsers", services.GetUsers)
    api.Post("/addUser",validations.AddUser, services.AddUser)
    api.Post("/updateUserById/:id",validations.EditUser, services.UpdateUser)
    api.Post("/deleteUserById/:id", services.DeleteUser)
}