package Model

type Course struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Name     string    `gorm:"unique" json:"name"`
	Credits  int       `json:"credits"`
	Code     string    `gorm:"unique" json:"code"`
	Classes  []Class   `gorm:"many2many:course_classes"`
	Students []Student `gorm:"many2many:course_student"`
}
