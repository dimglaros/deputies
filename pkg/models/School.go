package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Name       string
	Vacancies  []Vacancy
	DivisionID int
	Division   Division
}
