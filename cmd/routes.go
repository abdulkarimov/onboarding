package main

import (
	"github.com/abdulkarimov/onboarding/services"
	"github.com/abdulkarimov/onboarding/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/yandex"
)

func setupRoutes(app *fiber.App) {
	goth.UseProviders(
		google.New("1081271393484-28qa60cgsqeh34dcg1sbaqbde6m489lu.apps.googleusercontent.com", "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",
			"http://127.0.0.1:3000/api/auth/google/callback",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		),
		yandex.New("d2d27ae149c8410cbb58e017cf317df6",
			"aea23d1173d7429f8d533b58c853ab26",
			"http://127.0.0.1:3000/api/auth/yandex/callback",
		),
	)

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is working")
	})

	api.Get("/auth/:provider", services.Login)
	api.Get("/auth/:provider/callback", services.Callback)
	api.Get("/logout/:provider", services.Logout)

	api.Get("/getUsers", services.GetUsers)
	api.Post("/addUser", validations.AddUser, services.AddUser)
	api.Post("/updateUserById/:id", validations.EditUser, services.UpdateUser)
	api.Post("/deleteUserById/:id", services.DeleteUser)

}
