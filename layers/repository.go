package layers

import (
	"freepass-2023/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create() (models.NewUser, error)
	Read() (models.User, error)
	Update() (models.User, error)
	Delete() (models.User, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}
