package v1services

import (
	"context"

	"github.com/bagashiz/freepass-2023/db"
	"github.com/bagashiz/freepass-2023/models"
	"gorm.io/gorm"
)

// UpdateEnrollment is a struct for enrolling a student.
type UpdateEnrollmentParams struct {
	ClassCode  string `json:"class_code"`
	CourseCode string `json:"course_code"`
	StudentID  string `json:"student_id"`
}

// EnrollStudentToClass is a function that enrolls a student to a class.
func EnrollStudentToClass(ctx context.Context, arg UpdateEnrollmentParams) (models.Student, error) {
	var student models.Student

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get student.
		student, err = GetStudent(ctx, arg.StudentID)
		if err != nil {
			return err
		}

		// Get class.
		class, err := GetClass(ctx, arg.ClassCode)
		if err != nil {
			return err
		}

		// Get course.
		course, err := GetCourse(ctx, arg.CourseCode)
		if err != nil {
			return err
		}

		// Check if class is full.
		if class.Enrolled == class.Capacity {
			return gorm.ErrInvalidData
		}

		// Check if student total credits exceeds 24.
		if student.TotalCredits+course.Credits > 24 {
			return gorm.ErrInvalidData
		}

		// Check if student is already enrolled in the course.
		for _, class := range student.Classes {
			if class.CourseCode == course.CourseCode {
				return gorm.ErrInvalidData
			}
		}

		// Create enrollment.
		err = tx.Model(&class).Association("Students").Append(&student)
		if err != nil {
			return err
		}

		err = tx.Model(&student).Association("Classes").Append(&class)
		if err != nil {
			return err
		}

		// Update student total credits.
		student.TotalCredits += course.Credits
		err = tx.Save(&student).Error
		if err != nil {
			return err
		}

		// Update class enrolled.
		class.Enrolled++
		err = tx.Save(&class).Error
		if err != nil {
			return err
		}

		return nil
	})

	return student, err
}

// UnEnrollStudentFromClass is a function that unenrolls a student from a class.
func UnEnrollStudentFromClass(ctx context.Context, arg UpdateEnrollmentParams) (models.Student, error) {
	var student models.Student

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get student.
		student, err = GetStudent(ctx, arg.StudentID)
		if err != nil {
			return err
		}

		// Get class.
		class, err := GetClass(ctx, arg.ClassCode)
		if err != nil {
			return err
		}

		// Get course.
		course, err := GetCourse(ctx, arg.CourseCode)
		if err != nil {
			return err
		}

		// Check if student is not enrolled in the class.
		var enrolled bool
		for _, s := range class.Students {
			if s.StudentID == student.StudentID {
				enrolled = true
				break
			}
		}
		if !enrolled {
			return gorm.ErrInvalidData
		}

		// Delete enrollment.
		err = tx.Model(&class).Association("Students").Delete(&student)
		if err != nil {
			return err
		}

		err = tx.Model(&student).Association("Classes").Delete(&class)
		if err != nil {
			return err
		}

		// Update student total credits.
		student.TotalCredits -= course.Credits
		err = tx.Save(&student).Error
		if err != nil {
			return err
		}

		// Update class enrolled.
		class.Enrolled--
		err = tx.Save(&class).Error
		if err != nil {
			return err
		}

		return nil
	})

	return student, err
}
