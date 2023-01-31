package models

import "gorm.io/gorm"

// Course is a struct that represents a course entity.
type Course struct {
	CourseCode string  `gorm:"type:varchar(8);not null;unique;primaryKey" json:"course_code"`
	CourseName string  `gorm:"type:varchar;not null;unique" json:"course_name"`
	Credits    int     `gorm:"check:credits <= 6;check:credits >= 1" json:"credits"`
	Classes    []Class `gorm:"foreignKey:CourseCode;references:CourseCode" json:"classes"`
	gorm.Model
}
