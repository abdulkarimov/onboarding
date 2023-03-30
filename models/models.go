package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	DateOfBirth  string `json:"dateOfBirth"`
	Img          string `json:"img"`
	ProjectId    uint   `json:"projectId"`
	Project      Project
	StackId      uint `json:"stackId"`
	Stack        Stack
	PositionId   uint `json:"positionId"`
	Position     Position
	Info         string `json:"info"`
	DepartmentID uint   `json:"departmentId"`
	Department   Department
	ContactsID   uint `json:"contactsId"`
	Contacts     Contacts
	MentorID     uint `json:"mentorId"`
	Mentor       *User
	RoleID       uint `json:"roleId"`
	Role         Role
	Schedule     string `json:"schedule"`
	StatusID     uint   `json:"statusId"`
	Status       Status
	Verified     bool `json:"verified"`
}

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Stack struct {
	gorm.Model
	Name string `json:"name"`
}

type Position struct {
	gorm.Model
	Name string `json:"name"`
}

type Department struct {
	gorm.Model
	Name string `json:"name"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type Status struct {
	gorm.Model
	Name string `json:"name"`
}

type Contacts struct {
	gorm.Model
	HomePhone string `json:"homePhone"`
	WorkPhone string `json:"workPhone"`
	Email     string `json:"email"`
}

type Question struct {
	gorm.Model
	Name   string `json:"name"`
	FormId uint   `json:"formId"`
}

type Form struct {
	gorm.Model
	Name string `json:"name"`
}

type Post struct {
	gorm.Model
	Name    string `json:"name"`
	Content string `json:"content"`
	Img     string `json:"img"`
}

type Skill struct {
	gorm.Model
	Name string `json:"name"`
}

type Answer struct {
	gorm.Model
	Content    string `json:"content"`
	QuestionId uint   `json:"questionId"`
}
