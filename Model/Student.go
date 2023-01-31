package Model

type Student struct {
	ID       uint     `gorm:"primarykey" json:"id"`
	NIM      string   `gorm:"not null;unique" json:"nim"`
	Username string   `gorm:"not null;unique" json:"username"`
	Name     string   `gorm:"not null" json:"name"`
	Classes  []Class  `gorm:"many2many:student_classes"`
	Courses  []Course `gorm:"many2many:course_student"`
}
