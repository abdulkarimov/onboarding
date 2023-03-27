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
	api.Get("/getUsers", services.GetUsers)
	api.Post("/addUser", validations.AddUser, services.AddUser)
	api.Post("/updateUserById/:id", validations.EditUser, services.UpdateUser)
	api.Post("/deleteUserById/:id", services.DeleteUser)

	//project
	api.Get("/getProjects", repositories.GetProject)
	api.Post("/addProject", repositories.AddProject)
	api.Post("/updateProjectById/:id", repositories.UpdateProject)
	api.Post("/deleteProjectById/:id", repositories.DeleteProject)

	//stack
	api.Get("/getStacks", repositories.GetStacks)
	api.Post("/addStack", repositories.AddStack)
	api.Post("/updateStackById/:id", repositories.UpdateStack)
	api.Post("/deleteStackById/:id", repositories.DeleteStack)

	//position
	api.Get("/getPosition", repositories.GetPosition)
	api.Post("/addPosition", repositories.AddPosition)
	api.Post("/updatePositionById/:id", repositories.UpdatePosition)
	api.Post("/deletePositionById/:id", repositories.DeletePosition)

	//department
	api.Get("/getDepartment", repositories.GetDepartment)
	api.Post("/addDepartment", repositories.AddDepartment)
	api.Post("/updateDepartmentById/:id", repositories.UpdateDepartment)
	api.Post("/deleteDepartmentById/:id", repositories.DeleteDepartment)

	//role
	api.Get("/getRole", repositories.GetRole)
	api.Post("/addRole", repositories.AddRole)
	api.Post("/updateRoleById/:id", repositories.UpdateRole)
	api.Post("/deleteRoleById/:id", repositories.DeleteRole)

	//status
	api.Get("/getStatus", repositories.GetStatus)
	api.Post("/addStatus", repositories.AddStatus)
	api.Post("/updateStatusById/:id", repositories.UpdateStatus)
	api.Post("/deleteStatusById/:id", repositories.DeleteStatus)

	//contacts
	api.Get("/getContacts", repositories.GetContacts)
	api.Post("/addContacts", repositories.AddContacts)
	api.Post("/updateContactsById/:id", repositories.UpdateContacts)
	api.Post("/deleteContactsById/:id", repositories.DeleteContacts)

	//question
	api.Get("/getQuestion", repositories.GetQuestion)
	api.Post("/addQuestion", repositories.AddQuestion)
	api.Post("/updateQuestionById/:id", repositories.UpdateQuestion)
	api.Post("/deleteQuestionById/:id", repositories.DeleteQuestion)

	//form
	api.Get("/getForm", repositories.GetForm)
	api.Post("/addForm", repositories.AddForm)
	api.Post("/updateFormById/:id", repositories.UpdateForm)
	api.Post("/deleteFormById/:id", repositories.DeleteForm)

	//post
	api.Get("/getPost", repositories.GetPost)
	api.Post("/addPost", repositories.AddPost)
	api.Post("/updatePostById/:id", repositories.UpdatePost)
	api.Post("/deletePostById/:id", repositories.DeletePost)
}
