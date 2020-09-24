package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name         string
	Surname      string
	FatherName   string
	IdNumber     string
	Credentials  Credentials
	Applications []Application
}
