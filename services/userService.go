package services


import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/repositories"

)

func RegisterUser(c *fiber.Ctx) error {
	return repositories.RegisterUser(c)
}
func ShowIMG(c *fiber.Ctx) error {
	return repositories.ShowIMG(c)
}

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}

func GetUsers(c *fiber.Ctx) error {
	return repositories.GetUsers(c)
}

func AddUser(c *fiber.Ctx) error {
	return repositories.AddUser(c)
}

func UpdateUser( c *fiber.Ctx) error {
	return repositories.UpdateUser(c)
}

func DeleteUser( c *fiber.Ctx) error {
	return repositories.DeleteUser(c)
}
func GoogleAuth(c *fiber.Ctx) error {

	return repositories.GoogleAuth(c)
}

func GoogleCallback(c *fiber.Ctx) error {

	return repositories.GoogleCallback(c)
}