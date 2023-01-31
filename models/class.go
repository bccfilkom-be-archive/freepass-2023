package models

import "gorm.io/gorm"

type Class struct {
	ClassID   string `json:"class_id" gorm:"not null;unique;primaryKey"`
	ClassName string `json:"class_name" gorm:"not null;unique"`
	Credits   int    `json:"credits" gorm:"check:credits > 0;check:credits < 7"`
	Enrolled  int    `json:"enrolled" gorm:"check:enrolled <= capacity"`
	Students  []User `json:"students"`
	gorm.Model
}
