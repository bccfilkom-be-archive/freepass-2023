package v1services

import (
	"context"

	"github.com/bagashiz/freepass-2023/db"
	"github.com/bagashiz/freepass-2023/models"
	"gorm.io/gorm"
)

// RegisterStudentParams is a struct for registering a new student user.
type RegisterStudentParams struct {
	StudentID      string `json:"student_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// RegisterStudent is a function that registers a new student user.
func RegisterStudent(ctx context.Context, arg RegisterStudentParams) (models.Student, error) {
	student := models.Student{
		StudentID:      arg.StudentID,
		FullName:       arg.FullName,
		Email:          arg.Email,
		HashedPassword: arg.HashedPassword,
	}

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(&student).Error
	})

	return student, err
}

// GetStudent is a function that returns a student user based on StudentID.
func GetStudent(ctx context.Context, studentID string) (models.Student, error) {
	var student models.Student

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get student.
		if err = tx.First(&student, "student_id = ?", studentID).Error; err != nil {
			return err
		}

		// Get classes.
		if err = tx.Model(&student).Association("Classes").Find(&student.Classes); err != nil {
			return err
		}

		return nil
	})

	return student, err
}

// UpdateStudentParams is a struct for updating a student user.
type UpdateStudentParams struct {
	StudentID      string `json:"student_id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// UpdateStudent is a function that updates a student user.
func UpdateStudent(ctx context.Context, arg UpdateStudentParams) (models.Student, error) {
	var student models.Student

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		if student, err = GetStudent(ctx, arg.StudentID); err != nil {
			return err
		}

		student.Email = arg.Email
		student.HashedPassword = arg.HashedPassword

		return tx.Save(&student).Error
	})

	return student, err
}

// DeleteStudent is a function that deletes a student user.
func DeleteStudent(ctx context.Context, studentID string) (models.Student, error) {
	var student models.Student
	var classes []models.Class

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Check student if exists.
		if student, err = GetStudent(ctx, studentID); err != nil {
			return err
		}

		// Delete enrollments.
		if err = tx.Model(&student).Association("Classes").Find(&classes); err != nil {
			return err
		}

		for _, class := range classes {
			arg := UpdateEnrollmentParams{
				ClassCode:  class.ClassCode,
				CourseCode: class.CourseCode,
				StudentID:  student.StudentID,
			}

			if _, err = UnEnrollStudentFromClass(ctx, arg); err != nil {
				return err
			}
		}

		return tx.Delete(&student).Error
	})

	return student, err
}
