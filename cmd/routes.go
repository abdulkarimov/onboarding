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
	user.Post("/getById/:ID?", services.GetUsers)
	user.Post("/add", validations.AddUser, services.AddUser)
	user.Post("/update", validations.EditUser, services.UpdateUser)
	user.Post("/addProject", repositories.AddProjectToUser)
	user.Post("/addSkill", repositories.AddSkillToUser)
	user.Post("/addStack", repositories.AddStackToUser)
	user.Post("/delete/:ID?", services.DeleteUser)

	//project
	project := api.Group("/project")
	project.Get("/get", repositories.GetProjects)
	project.Post("/add", repositories.AddProject)
	project.Post("/update", repositories.UpdateProject)
	project.Post("/delete/:ID?", repositories.DeleteProject)

	//stack
	stack := api.Group("/stack")
	stack.Get("/get", repositories.GetStacks)
	stack.Post("/add", repositories.AddStack)
	stack.Post("/update", repositories.UpdateStack)
	stack.Post("/delete/:ID?", repositories.DeleteStack)

	//position
	position := api.Group("/position")
	position.Get("/get", repositories.GetPositions)
	position.Post("/add", repositories.AddPosition)
	position.Post("/update", repositories.UpdatePosition)
	position.Post("/delete/:ID?", repositories.DeletePosition)

	//department
	department := api.Group("/department")
	department.Get("/get", repositories.GetDepartments)
	department.Post("/add", repositories.AddDepartment)
	department.Post("/update", repositories.UpdateDepartment)
	department.Post("/delete/:ID?", repositories.DeleteDepartment)

	//role
	role := api.Group("/role")
	role.Get("/get", repositories.GetRoles)
	role.Post("/add", repositories.AddRole)
	role.Post("/update", repositories.UpdateRole)
	role.Post("/delete/:ID?", repositories.DeleteRole)

	//status
	status := api.Group("/status")
	status.Get("/get", repositories.GetStatuses)
	status.Post("/add", repositories.AddStatus)
	status.Post("/update", repositories.UpdateStatus)
	status.Post("/delete/:ID?", repositories.DeleteStatus)

	//contacts
	contacts := api.Group("/contacts")
	contacts.Get("/get", repositories.GetContacts)
	contacts.Post("/add", repositories.AddContacts)
	contacts.Post("/update", repositories.UpdateContacts)
	contacts.Post("/delete/:ID?", repositories.DeleteContacts)

	//question
	question := api.Group("/question")
	question.Get("/get", repositories.GetQuestions)
	question.Post("/add", repositories.AddQuestion)
	question.Post("/update", repositories.UpdateQuestion)
	question.Post("/delete/:ID?", repositories.DeleteQuestion)

	//form
	form := api.Group("/form")
	form.Get("/get", repositories.GetForms)
	form.Post("/add", repositories.AddForm)
	form.Post("/update", repositories.UpdateForm)
	form.Post("/addQuestion", repositories.AddQuestionToForm)
	form.Post("/delete/:ID?", repositories.DeleteForm)

	//post
	post := api.Group("/post")
	post.Get("/get", repositories.GetPosts)
	post.Post("/add", repositories.AddPost)
	post.Post("/update", repositories.UpdatePost)
	post.Post("/delete/:ID?", repositories.DeletePost)

	//skill
	skill := api.Group("/skill")
	skill.Get("/get", repositories.GetSkills)
	skill.Post("/add", repositories.AddSkill)
	skill.Post("/update", repositories.UpdateSkill)
	skill.Post("/delete/:ID?", repositories.DeleteSkill)

	//answer
	answer := api.Group("/answer")
	answer.Get("/get", repositories.GetAnswers)
	answer.Post("/add", repositories.AddAnswer)
	answer.Post("/update", repositories.UpdateAnswer)
	answer.Post("/addQuestion", repositories.AddAnswerToQuestion)
	answer.Post("/delete/:ID?", repositories.DeleteAnswer)

}
