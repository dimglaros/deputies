package db

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name       string
	Surname    string
	FatherName string
	IdNumber   string
	Email      string
	Applications []Application
}

type Application struct {
	gorm.Model
	Year uint
	SpecialTreatment bool
	ZEP bool
	TeacherID int
	Teacher Teacher
}
