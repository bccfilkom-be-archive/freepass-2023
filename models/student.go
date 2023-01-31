package models

import "gorm.io/gorm"

// Student is a struct that represents a student entity.
type Student struct {
	StudentID      string  `gorm:"type:varchar;not null;unique;primaryKey" json:"student_id"`
	FullName       string  `gorm:"type:varchar;not null" json:"full_name"`
	Email          string  `gorm:"type:varchar;not null;unique" json:"email"`
	HashedPassword string  `gorm:"type:varchar;not null" json:"hashed_password"`
	TotalCredits   int     `gorm:"check:total_credits <= 24;check:total_credits >= 0" json:"total_credits"`
	Classes        []Class `gorm:"many2many:enrollments" json:"classes"`
	gorm.Model
}
