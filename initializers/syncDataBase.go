package initializers

import "freepass-2023/models"

func SyncDataBase() {
	DB.AutoMigrate(&models.NewUser{}, &models.User{})
}
