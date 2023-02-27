package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/repositories"
    "github.com/abdulkarimov/onboarding/validations"
)

func setupRoutes(app *fiber.App) {
    api := app.Group("/api")  

    api.Get("/getUsers", repositories.GetUsers)
    api.Post("/addUser",validations.AddUser, repositories.AddUser)
    api.Post("/updateUserById/:id",validations.EditUser, repositories.UpdateUser)
    api.Post("/deleteUserById/:id", repositories.DeleteUser)

}