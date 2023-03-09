package main

import (
	"github.com/abdulkarimov/onboarding/services"
	"github.com/abdulkarimov/onboarding/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/yandex"
	"os"
)

func setupRoutes(app *fiber.App) {
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

	api := app.Group("/api")

	api.Get("/getUsers", services.GetUsers)
	api.Post("/addUser", validations.AddUser, services.AddUser)
	api.Post("/updateUserById/:id", validations.EditUser, services.UpdateUser)
	api.Post("/deleteUserById/:id", services.DeleteUser)

	auth := app.Group("/auth")
	auth.Get("/:provider", services.Login)
	auth.Get("/:provider/callback", services.Callback)
	auth.Get("/logout/:provider", services.Logout)
	auth.Get("/user/verify", services.Verify)
}
