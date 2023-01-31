package models

import "gorm.io/gorm"

type Course struct {
	CourseName string `json:"course_name" gorm:"not null;unique;primaryKey"`
	Credits    int    `json:"credits" gorm:"check:credits > 0;check:credits < 7"`
	gorm.Model
}
