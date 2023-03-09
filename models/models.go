package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name string  
    Age int 
    Email string 
    StatusId int
    Status Status
}

type Status struct {
    Id int
    Name string
}