package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Score            float32
	Year             uint
	SpecialTreatment bool
	ZEP              bool
	TeacherID        int
}
