package models

import "gorm.io/gorm"

// Class is a struct that represents a class entity.
type Class struct {
	ClassCode  string    `gorm:"type:varchar(10);not null;unique;primaryKey" json:"class_code"`
	CourseCode string    `gorm:"type:varchar(8);not null" json:"course_code"`
	Capacity   int       `gorm:"check:capacity <= 40;check:capacity >= 20" json:"capacity"`
	Enrolled   int       `gorm:"check:enrolled <= capacity" json:"enrolled"`
	Students   []Student `gorm:"many2many:enrollments" json:"students"`
	gorm.Model
}
