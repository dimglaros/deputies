package models

import "gorm.io/gorm"

type Division struct {
	gorm.Model
	Name string
}
