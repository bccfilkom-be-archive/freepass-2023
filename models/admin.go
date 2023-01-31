package models

type Admin struct {
	AdminID  string `json:"admin_id" gorm:"primarykey;unique;not null"`
	Fullname string `json:"fullname" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
