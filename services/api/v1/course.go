package v1services

import (
	"context"

	"github.com/bagashiz/freepass-2023/db"
	"github.com/bagashiz/freepass-2023/models"
	"gorm.io/gorm"
)

// RegisterCourseParams is a struct for registering a new course.
type RegisterCourseParams struct {
	CourseCode string `json:"course_code"`
	CourseName string `json:"course_name"`
	Credits    int    `json:"credits"`
}

// RegisterCourse is a function that registers a new course.
func RegisterCourse(ctx context.Context, arg RegisterCourseParams) (models.Course, error) {
	course := models.Course{
		CourseCode: arg.CourseCode,
		CourseName: arg.CourseName,
		Credits:    arg.Credits,
	}

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(&course).Error
	})

	return course, err
}

// GetCourse is a function that returns a course based on course code.
func GetCourse(ctx context.Context, courseCode string) (models.Course, error) {
	var course models.Course

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get course.
		if err = tx.First(&course, "course_code = ?", courseCode).Error; err != nil {
			return err
		}

		// Get classes with the same course code.
		if err = tx.Where("course_code = ?", courseCode).Find(&course.Classes).Error; err != nil {
			return err
		}

		return nil
	})

	return course, err
}

// ListCoursesParams is a struct for listing courses with pagination.
type ListCoursesParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// ListCourses is a function that returns a list of courses with pagination and sorted by course code.
func ListCourses(ctx context.Context, arg ListCoursesParams) ([]models.Course, error) {
	var courses []models.Course

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Get courses.
		if err = tx.Limit(arg.Limit).Offset(arg.Offset).Order("course_code").Find(&courses).Error; err != nil {
			return err
		}

		// Get classes for each course.
		for i := range courses {
			if err = tx.Where("course_code = ?", courses[i].CourseCode).Find(&courses[i].Classes).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return courses, err
}

// UpdateCourseParams is a struct for updating a course.
type UpdateCourseParams struct {
	CourseCode string `json:"course_code"`
	CourseName string `json:"course_name"`
	Credits    int    `json:"credits"`
}

// UpdateCourse is a function that updates a course.
func UpdateCourse(ctx context.Context, arg UpdateCourseParams) (models.Course, error) {
	var course models.Course

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		// Check if course exists.
		if course, err = GetCourse(ctx, arg.CourseCode); err != nil {
			return err
		}

		// Check if course already has class(es).
		if len(course.Classes) > 0 {
			return gorm.ErrInvalidData
		}

		course.CourseName = arg.CourseName
		course.Credits = arg.Credits

		return tx.Save(&course).Error
	})

	return course, err
}

// DeleteCourse is a function that deletes a course.
func DeleteCourse(ctx context.Context, courseCode string) (models.Course, error) {
	var course models.Course

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		if course, err = GetCourse(ctx, courseCode); err != nil {
			return err
		}

		return tx.Unscoped().Delete(&course).Error
	})

	return course, err
}
