package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50);unique" json:"email"`
	Username string `gorm:"type:varchar(255);unique" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}

type Users struct {
	gorm.Model
	Nim string `gorm:"type:varchar(50);unique" json:"nim"`
	Email string `gorm:"type:varchar(50);unique" json:"email"`
	Username string `gorm:"type:varchar(255);unique" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Sks int `gorm:"type:int" json:"sks"`
}

type Courses struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);unique" json:"title"`
	Course_code string `gorm:"type:varchar(255);unique" json:"course_code"`
	Sks int `gorm:"type:int" json:"sks"`
}


type Class struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique" json:"name"`
	Class_code string `gorm:"type:varchar(255);unique" json:"class_code"`
	Location string `gorm:"type:varchar(255)" json:"location"`
	CourseID uint
	Course Courses `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserClass struct{
	UserID uint
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClassID uint
	Class Class `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
