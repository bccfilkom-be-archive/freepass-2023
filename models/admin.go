package models

import "gorm.io/gorm"

// Admin is a struct that represents an admin entity.
type Admin struct {
	AdminID        string `gorm:"type:varchar;not null;unique;primaryKey" json:"admin_id"`
	FullName       string `gorm:"type:varchar;not null" json:"full_name"`
	Email          string `gorm:"type:varchar;not null;unique" json:"email"`
	HashedPassword string `gorm:"type:varchar;not null" json:"hashed_password"`
	gorm.Model
}
