package db

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name         string
	Surname      string
	FatherName   string
	IdNumber     string
	Email        string
	Applications []Application
}

type Application struct {
	gorm.Model
	Score            float32
	Year             uint
	SpecialTreatment bool
	ZEP              bool
	TeacherID        int
	Divisions        []*Division `gorm:"many2many:application_divisions;"`
}

type Vacancy struct {
	gorm.Model
	Year      uint
	Specialty string
	SchoolID  int
}

type School struct {
	gorm.Model
	Name       string
	Vacancies  []Vacancy
	DivisionID int
	Division   Division
}

type Division struct {
	gorm.Model
	Name         string
	Applications []*Application `gorm:"many2many:application_divisions;"`
}
