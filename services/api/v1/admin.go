package v1services

import (
	"context"

	"github.com/bagashiz/freepass-2023/db"
	"github.com/bagashiz/freepass-2023/models"
	"gorm.io/gorm"
)

// RegisterAdminParams is a struct for registering a new admin user.
type RegisterAdminParams struct {
	AdminID        string `json:"admin_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// RegisterAdmin is a function that registers a new admin user.
func RegisterAdmin(ctx context.Context, arg RegisterAdminParams) (models.Admin, error) {
	admin := models.Admin{
		AdminID:        arg.AdminID,
		FullName:       arg.FullName,
		Email:          arg.Email,
		HashedPassword: arg.HashedPassword,
	}

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(&admin).Error
	})

	return admin, err
}

// GetAdmin is a function that returns an admin user based on AdminID.
func GetAdmin(ctx context.Context, adminID string) (models.Admin, error) {
	var admin models.Admin

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.First(&admin, "admin_id = ?", adminID).Error
	})

	return admin, err
}

// UpdateAdminParams is a struct for updating an admin user.
type UpdateAdminParams struct {
	AdminID        string `json:"admin_id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// UpdateAdmin is a function that updates an admin user.
func UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (models.Admin, error) {
	var admin models.Admin

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		if admin, err = GetAdmin(ctx, arg.AdminID); err != nil {
			return err
		}

		admin.Email = arg.Email
		admin.HashedPassword = arg.HashedPassword

		return tx.Save(&admin).Error
	})

	return admin, err
}

// DeleteAdmin is a function that deletes an admin user.
func DeleteAdmin(ctx context.Context, adminID string) (models.Admin, error) {
	var admin models.Admin

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		if admin, err = GetAdmin(ctx, adminID); err != nil {
			return err
		}

		return tx.Delete(&admin).Error
	})

	return admin, err
}
