package models

import "gorm.io/gorm"

type Class struct {
	ClassName string   `json:"class_name" gorm:"not null;unique;primaryKey"`
	Courses   []Course `json:"courses"`
	Enrolled  int      `json:"enrolled" gorm:"check:enrolled <= capacity"`
	Students  []User   `json:"students"`
	gorm.Model
}
