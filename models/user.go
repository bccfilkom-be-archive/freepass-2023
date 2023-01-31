package models

type User struct {
	StudentID     string   `json:"student_id" gorm:"primarykey;unique;not null"`
	Fullname      string   `json:"fullname" gorm:"not null"`
	Email         string   `json:"email" gorm:"unique;not null"`
	Password      string   `json:"password" gorm:"not null"`
	SKS           int      `json:"sks" gorm:"not null;check:sks >= 0;check:sks <= 24" `
	EnrolledClass []string `json:"enrolled_class"`
}
