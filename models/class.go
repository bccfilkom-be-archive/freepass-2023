package models

import "gorm.io/gorm"

// Class is a struct that represents a class entity.
type Class struct {
	ClassName string   `json:"class_name" gorm:"not null;unique;primaryKey"`
	Courses   []Course `json:"courses"`
	Enrolled  int      `json:"enrolled" gorm:"check:enrolled <= capacity"`
	Students  []User   `json:"students"`
	gorm.Model
}
