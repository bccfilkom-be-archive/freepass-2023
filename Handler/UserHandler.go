package Handler

import (
	"freepass/Middleware"
	"freepass/Model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	// Get user profile
	r.GET("/profile", Middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.Role != 0 && user.Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "You are not an admin or a student",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"data": user.Username,
		})
	})

	// Edit user profile
	r.POST("/profile", Middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.Role != 0 && user.Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "You are not an admin or a student",
			})
		}

		type inputEditUser struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			NIM      string `json:"nim"`
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var input inputEditUser

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Error when binding JSON",
				"error":   err.Error(),
			})
			return
		}

		if input.Name == "" || input.NIM == "" || input.Username == "" || input.Password == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Cannot have any empty field.",
			})
			return
		}

		updatedUser := Model.User{
			Email:     input.Email,
			Username:  input.Username,
			Password:  Hash(input.Password),
			UpdatedAt: time.Now(),
		}

		if err := db.Where("id = ?", ID).Model(&updatedUser).Updates(updatedUser); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error updating data user",
				"error":   err.Error.Error(),
			})
		}

		updatedStudent := Model.Student{
			Username: input.Username,
			Name:     input.Name,
			NIM:      input.NIM,
		}

		if err := db.Where("id = ?", ID).Model(&updatedStudent).Updates(updatedStudent); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error updating data student",
				"error":   err.Error.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"message": "successfully edited.",
			// "data":    updatedUser,
		})
	})

	// Get all class
	r.GET("/class", Middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.Role != 0 && user.Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		var classes []Model.Class

		if err := db.Preload("Courses").Find(&classes); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong when querrying",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"data": classes,
		})
	})

	// Add user to class
	r.POST("/class", Middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.Role != 0 && user.Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		type inputJSON struct {
			ClassID uint `json:"class_id"`
		}

		var input inputJSON
		var class Model.Class
		var student Model.Student

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot bind JSON class.",
				"error":   err.Error(),
			})
		}

		if err := db.Where("id = ?", input.ClassID).First(&class).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Class not found",
				"error":   err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", user.ID).First(&student).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Student not found",
				"error":   err.Error(),
			})
			return
		}

		var studentExists bool
		query := `SELECT EXISTS(SELECT 1 FROM student_classes WHERE student_id = ? AND class_id = ?)`
		if err := db.Raw(query, user.ID, input.ClassID).Row().Scan(&studentExists); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Theres already that student in that class",
				"error":   err.Error(),
			})
			return
		}

		if studentExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Theres already that student in that class",
			})
			return
		}

		if err := db.Where("id = ?", user.ID).Preload("Courses").First(&student).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "students course not found",
				"error":   err.Error(),
			})
			return
		}

		totalCredit := 0

		for _, course := range student.Courses {
			totalCredit += course.Credits
		}

		var studentCourseExists bool
		query = `SELECT EXISTS(SELECT 1 FROM student_classes WHERE student_id = ? AND class_id = ?)`
		if err := db.Raw(query, user.ID, input.ClassID).Row().Scan(&studentExists); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Theres already that student in that class",
				"error":   err.Error(),
			})
			return
		}

		if studentCourseExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Theres already that student in that class",
			})
			return
		}

		var course Model.Course
		if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN course_classes ON courses.id = course_classes.course_id").Where("course_classes.class_id = ?", input.ClassID).First(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "failed to fetch course",
				"error":   err.Error(),
			})
			return
		}

		var studentCourse bool
		query = `SELECT EXISTS(SELECT 1 FROM course_student WHERE student_id = ? AND course_id = ?)`
		if err := db.Raw(query, user.ID, course.ID).Row().Scan(&studentCourse); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Theres already that student in that class",
				"error":   err.Error(),
			})
			return
		}

		if studentCourse {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Student already took that course in the other class",
			})
			return
		}

		if totalCredit+course.Credits > 24 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Student cannot take more than 24 credits",
			})
			return
		}

		if err := db.Model(&class).Association("Students").Append(&student); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot add student to class",
				"error":   err.Error(),
			})
			return
		}

		if err := db.Model(&course).Association("Students").Append(&student); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot add student to course",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"message": "successfully added new class",
			// "data":    student,
		})
	})

	// Remove user from class
	r.POST("/removeFromClass", Middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.Role != 0 && user.Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		type inputJSON struct {
			ClassID uint `json:"class_id"`
		}

		var input inputJSON

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot bind JSON.",
				"error":   err.Error(),
			})
			return
		}

		if err := db.Exec("DELETE FROM student_classes WHERE student_id = ? AND class_id = ?", user.ID, input.ClassID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot delete student from class",
				"error":   err.Error(),
			})
			return
		}

		var delCourse Model.Course
		if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN course_classes ON courses.id = course_classes.course_id").Where("course_classes.class_id = ?", input.ClassID).First(&delCourse).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "failed to fetch course",
				"error":   err.Error(),
			})
			return
		}

		if err := db.Exec("DELETE FROM course_student WHERE student_id = ? AND course_id = ?", user.ID, delCourse.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot delete student from course",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"message": "successfully dropped a class.",
		})
	})

	// Get all classes and participant
	r.GET("/classes", Middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.Role != 0 && user.Role != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		var student Model.Student
		if err := db.Where("id = ?", ID).Take(&student); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		type StudentData struct {
			Name string `json:"name"`
		}

		type ClassData struct {
			Name     string        `json:"name"`
			Students []StudentData `json:"students"`
		}

		var classes []Model.Class
		var studentData []StudentData
		var classData []ClassData

		if err := db.Model(&student).Association("Classes").Find(&classes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Class not found",
				"error":   err.Error(),
			})
			return
		}

		for _, class := range classes {
			studentData = nil
			if err := db.Model(class).Association("Students").Find(&studentData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "student not found",
					"success": false,
				})
				return
			}

			classData = append(classData, ClassData{
				Name:     class.Name,
				Students: studentData,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"class": classData,
		})
	})
}
