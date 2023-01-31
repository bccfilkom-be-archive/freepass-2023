package Model

import "time"

type User struct {
	ID        uint    `gorm:"primarykey" json:"id"`
	Email     string  `gorm:"not null;unique" json:"email"`
	Username  string  `gorm:"not null;unique" json:"username"`
	Password  string  `gorm:"not null" json:"password"`
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"ForeignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// role = 0 => admin
	// role = 1 => student
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
