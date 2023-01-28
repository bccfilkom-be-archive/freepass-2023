package initializers

import "course-management/models"

func Syncdatabase() {
	DB.AutoMigrate(
		&models.Users{},
		&models.Admin{},
		&models.Courses{},
		&models.Class{},
		&models.UserClass{},
	)
}