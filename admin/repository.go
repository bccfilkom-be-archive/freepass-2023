package admin

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Admin, error)
	FindByID(ID int) (Admin, error)
	Create(admin Admin) (Admin, error)
	Update(admin Admin) (Admin, error)
	Delete(admin Admin) (Admin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Admin, error) {
	var admins []Admin
	err := r.db.Find(&admins).Error

	return admins, err
}

func (r *repository) FindByID(ID int) (Admin, error) {
	var admin Admin

	err := r.db.Find(&admin, ID).Error

	return admin, err
}

func (r *repository) Create(admin Admin) (Admin, error) {

	err := r.db.Create(&admin).Error

	return admin, err
}

func (r *repository) Update(admin Admin) (Admin, error) {
	err := r.db.Save(&admin).Error

	return admin, err
}

func (r *repository) Delete(admin Admin) (Admin, error) {
	err := r.db.Delete(&admin).Error

	return admin, err
}
