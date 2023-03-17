package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name          string
	Age           int
	Date_of_birth time.Time
	Img           string
	PositionId    uint
	Position      Position
	Info          string
	DepartmentID  uint
	Department    Department
	CabinetID     uint
	Cabinet       Cabinet
	ContactsID    uint
	Contacts      Contacts
	MentorID      uint
	Mentor        *User
	Questions     []Question `gorm:"many2many:userQuestion"`
	Stacks        []Stack    `gorm:"many2many:userStack"`
	Projects      []Project  `gorm:"many2many:userProject"`
	HashID        uint
	Hash          Hash
	RoleID        uint
	Role          Role
	Schedule      string
	StatusID      uint
	Status        Status
}

type Project struct {
	gorm.Model
	Name        string
	Description string
}

type Stack struct {
	ID   uint `gorm:"primary key"`
	Name string
}

type Position struct {
	ID   uint `gorm:"primary key"`
	Name string
}

type Department struct {
	ID   uint `gorm:"primary key"`
	Name string
}

type Role struct {
	ID   uint `gorm:"primary key"`
	Name string
}

type Status struct {
	ID   uint `gorm:"primary key"`
	Name string
}

type Hash struct {
	ID   uint `gorm:"primary key"`
	Name string
}

type Cabinet struct {
	ID      uint `gorm:"primary key"`
	Floor   uint
	Cabinet uint
}

type Contacts struct {
	ID         uint `gorm:"primary key"`
	Home_phone string
	Work_phone string
	Email      string
}

type Question struct {
	gorm.Model
	Name   string
	FormId uint
	Form   Form
}

type Form struct {
	ID   uint `gorm:"primary key"`
	Name string
}
