package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Surname    string
	FatherName string
	IdNumber   string
	Email      string
}
