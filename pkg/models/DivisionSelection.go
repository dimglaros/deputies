package models

import "gorm.io/gorm"

type DivisionSelection struct {
	gorm.Model
	order         uint
	ApplicationID uint
	DivisionID    uint
	Division      Division
}
