package main

import (
	"github.com/abdulkarimov/onboarding/repositories"
	"github.com/abdulkarimov/onboarding/services"
	"github.com/abdulkarimov/onboarding/validations"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	//user
	user := api.Group("/user")
	user.Get("/get", services.GetUsers)
	user.Post("/add", validations.AddUser, services.AddUser)
	user.Post("/update", validations.EditUser, services.UpdateUser)
	user.Post("/delete", services.DeleteUser)

	//project
	project := api.Group("/project")
	project.Get("/get", repositories.GetProject)
	project.Post("/add", repositories.AddProject)
	project.Post("/update", repositories.UpdateProject)
	project.Post("/delete", repositories.DeleteProject)

	//stack
	stack := api.Group("/stack")
	stack.Get("/get", repositories.GetStacks)
	stack.Post("/add", repositories.AddStack)
	stack.Post("/update", repositories.UpdateStack)
	stack.Post("/delete", repositories.DeleteStack)

	//position
	position := api.Group("/position")
	position.Get("/get", repositories.GetPosition)
	position.Post("/add", repositories.AddPosition)
	position.Post("/update", repositories.UpdatePosition)
	position.Post("/delete", repositories.DeletePosition)

	//department
	department := api.Group("/department")
	department.Get("/get", repositories.GetDepartment)
	department.Post("/add", repositories.AddDepartment)
	department.Post("/update", repositories.UpdateDepartment)
	department.Post("/delete", repositories.DeleteDepartment)

	//role
	role := api.Group("/role")
	role.Get("/get", repositories.GetRole)
	role.Post("/add", repositories.AddRole)
	role.Post("/update", repositories.UpdateRole)
	role.Post("/delete", repositories.DeleteRole)

	//status
	status := api.Group("/status")
	status.Get("/get", repositories.GetStatus)
	status.Post("/add", repositories.AddStatus)
	status.Post("/update", repositories.UpdateStatus)
	status.Post("/delete", repositories.DeleteStatus)

	//contacts
	contacts := api.Group("/contacts")
	contacts.Get("/get", repositories.GetContacts)
	contacts.Post("/add", repositories.AddContacts)
	contacts.Post("/update", repositories.UpdateContacts)
	contacts.Post("/delete", repositories.DeleteContacts)

	//question
	question := api.Group("/question")
	question.Get("/get", repositories.GetQuestion)
	question.Post("/add", repositories.AddQuestion)
	question.Post("/update", repositories.UpdateQuestion)
	question.Post("/delete", repositories.DeleteQuestion)

	//form
	form := api.Group("/form")
	form.Get("/get", repositories.GetForm)
	form.Post("/add", repositories.AddForm)
	form.Post("/update", repositories.UpdateForm)
	form.Post("/delete", repositories.DeleteForm)

	//post
	post := api.Group("/post")
	post.Get("/get", repositories.GetPost)
	post.Post("/add", repositories.AddPost)
	post.Post("/update", repositories.UpdatePost)
	post.Post("/delete", repositories.DeletePost)
}
