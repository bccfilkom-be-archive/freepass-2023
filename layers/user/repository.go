package layers

import (
	"freepass-2023/models"

	"gorm.io/gorm"
)

type Repository interface {
	AddClass(ID int) (models.User, error)
	ViewClass() ([]models.Class, error)
	ViewParticipants() (models.Class, error)
	EditProfile() (models.User, error)
	DropClass() (models.User, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) AddClass(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repo) ViewClass() ([]models.Class, error) {
	var class []models.Class
	err := r.db.Find(&class).Error
	return class, err
}

func (r *repo) ViewParticipants() (models.Class, error) {
	var class models.Class
	err := r.db.Select("students").Find(&class).Error
	return class, err
}

