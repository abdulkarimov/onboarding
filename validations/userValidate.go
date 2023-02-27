package validations

import (
    "github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"
)
type IError struct {
    Field string
    Tag   string
    Value string
}

var Validator = validator.New()

func AddUser(c *fiber.Ctx) error {

    test := struct {
        Name  string `validate:"required"`
        Age int `validate:"required,number"`
    }{}

    c.BodyParser(&test)

    err := Validator.Struct(test)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(validate(err))
    }

    return c.Next()
}

func EditUser(c *fiber.Ctx) error {

    test := struct {
        Name  string 
        Age int 
    }{}
    c.BodyParser(&test)

    err := Validator.Struct(test)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(validate(err))
    }

    return c.Next()
}


func validate (err error ) []*IError {
    var errors []*IError
    for _, err := range err.(validator.ValidationErrors) {
        var el IError
        el.Field = err.Field()
        el.Tag = err.Tag()
        el.Value = err.Param()
        errors = append(errors, &el)
    }

    return errors
}