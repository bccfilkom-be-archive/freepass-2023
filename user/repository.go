package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *repository) Create(user User) (User, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) Delete(user User) (User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
