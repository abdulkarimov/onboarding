package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/services"
    "github.com/abdulkarimov/onboarding/validations"
    "github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
    "os"
	// "github.com/shareed2k/goth_fiber"
    // "log"
)

func setupRoutes(app *fiber.App) {

    goth.UseProviders(
        google.New(os.Getenv("OAUTH_KEY"), os.Getenv("OAUTH_SECRET"), "http://127.0.0.1:3000/api/auth/google/callback"),
    )
  
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{})
    })

    api := app.Group("/api")  

    api.Get("/next", services.ShowIMG)
    api.Get("/home", services.Home)
    api.Get("/registerUser/:email", services.RegisterUser)
    api.Get("/getUsers", services.GetUsers)
    api.Post("/addUser",validations.AddUser, services.AddUser)
    api.Post("/updateUserById/:id",validations.EditUser, services.UpdateUser)
    api.Post("/deleteUserById/:id", services.DeleteUser)
    api.Get("/auth/:provider", services.GoogleAuth)
    api.Get("/auth/:provider/callback", services.GoogleCallback)

	// api.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
	// 	gf.Logout(ctx)
	// 	ctx.Redirect("/")
	// 	return nil
	// })


}