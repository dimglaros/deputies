package models

import "gorm.io/gorm"

type Vacancy struct {
	gorm.Model
	Year      uint
	Specialty string
	SchoolID  int
}
