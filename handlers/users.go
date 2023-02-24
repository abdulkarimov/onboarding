package handlers


import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/database"
    "github.com/abdulkarimov/onboarding/models"
)



func GetUsers(c *fiber.Ctx) error {
    user := []models.User{}
    database.DB.Db.Find(&user)

    return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
        user := new(models.User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": err.Error(),
            })
        }
    
        database.DB.Db.Create(&user)
    
        return c.Status(200).JSON(user)
}