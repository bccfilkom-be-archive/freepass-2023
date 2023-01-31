package v1services

import (
	"context"

	"github.com/bagashiz/freepass-2023/db"
	"github.com/bagashiz/freepass-2023/models"
	"gorm.io/gorm"
)

// RegisterClassParams is a struct for registering a new class.
type RegisterClassParams struct {
	ClassCode  string `json:"class_code"`
	CourseCode string `json:"course_code"`
	Capacity   int    `json:"capacity"`
}

// RegisterClass is a function that registers a new class.
func RegisterClass(ctx context.Context, arg RegisterClassParams) (models.Class, error) {
	class := models.Class{}

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get course to check if it exists.
		course, err := GetCourse(ctx, arg.CourseCode)
		if err != nil {
			return err
		}

		// Create new class.
		class = models.Class{
			ClassCode:  arg.ClassCode,
			CourseCode: course.CourseCode,
			Capacity:   arg.Capacity,
			Enrolled:   0,
		}

		return tx.Create(&class).Error
	})

	return class, err
}

// GetClass is a function that returns a class based on class code.
func GetClass(ctx context.Context, classCode string) (models.Class, error) {
	var class models.Class

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get class.
		if err = tx.First(&class, "class_code = ?", classCode).Error; err != nil {
			return err
		}

		// Get students.
		if err = tx.Model(&class).Association("Students").Find(&class.Students); err != nil {
			return err
		}

		return nil
	})

	return class, err
}

// ListClasesParams is a struct for listing classes with pagination.
type ListClasesParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// ListClasses is a function that returns a list of classes with pagination and sorted by code.
func ListClasses(ctx context.Context, arg ListClasesParams) ([]models.Class, error) {
	var classes []models.Class

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get classes.
		if err = tx.Limit(arg.Limit).Offset(arg.Offset).Order("class_code").Find(&classes).Error; err != nil {
			return err
		}

		// Get students for each class.
		for i := range classes {
			if err = tx.Model(&classes[i]).Association("Students").Find(&classes[i].Students); err != nil {
				return err
			}
		}

		return nil
	})

	return classes, err
}

// UpdateClassParams is a struct for updating a class.
type UpdateClassParams struct {
	ClassCode  string `json:"class_code"`
	CourseCode string `json:"course_code"`
	Capacity   int    `json:"capacity"`
}

// UpdateClass is a function that updates a class.
func UpdateClass(ctx context.Context, arg UpdateClassParams) (models.Class, error) {
	var class models.Class

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		if class, err = GetClass(ctx, arg.ClassCode); err != nil {
			return err
		}

		// Get course to check if it exists.
		course, err := GetCourse(ctx, arg.CourseCode)
		if err != nil {
			return err
		}

		// New capacity must be greater than enrolled.
		if arg.Capacity < class.Enrolled {
			return gorm.ErrInvalidData
		}

		class.CourseCode = course.CourseCode
		class.Capacity = arg.Capacity

		return tx.Save(&class).Error
	})

	return class, err
}

// DeleteClass is a function that deletes a class based on class code.
func DeleteClass(ctx context.Context, classCode string) (models.Class, error) {
	var class models.Class
	var students []models.Student

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Check if class exists.
		if class, err = GetClass(ctx, classCode); err != nil {
			return err
		}

		// Delete enrollments.
		if err = tx.Model(&class).Association("Students").Find(&students); err != nil {
			return err
		}

		for _, student := range students {
			arg := UpdateEnrollmentParams{
				ClassCode:  class.ClassCode,
				CourseCode: class.CourseCode,
				StudentID:  student.StudentID,
			}

			if _, err = UnEnrollStudentFromClass(ctx, arg); err != nil {
				return err
			}
		}

		return tx.Unscoped().Delete(&class).Error
	})

	return class, err
}
