package Model

type Admin struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}
