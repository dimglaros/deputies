package models

import "gorm.io/gorm"

type Credentials struct {
	gorm.Model
	Email     string
	Password  string
	TeacherID uint
}
