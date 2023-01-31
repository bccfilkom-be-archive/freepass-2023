package Model

import "time"

type Class struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    ` json:"name"`
	Students  []Student `gorm:"many2many:student_classes"`
	Courses   []Course  `gorm:"many2many:course_classes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
