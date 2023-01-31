package Handler

import (
	"freepass/Middleware"
	"freepass/Model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Admin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/admin")
	// view all class
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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}

		var classes []Model.Class
		if err := db.Preload("Courses", "name").Find(&classes); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot get classes.",
			})
			return
		}

		type classResult struct {
			ID   uint
			Name string
		}

		var data []classResult
		for _, value := range classes {
			var temp classResult
			temp.ID = value.ID
			temp.Name = value.Name
			data = append(data, temp)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// show all classes with students and course
	r.GET("/classDetailed", Middleware.Authorization(), func(c *gin.Context) {
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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}

		type StudentData struct {
			Name string `json:"name"`
		}

		type CourseData struct {
			Name    string `json:"name"`
			Credits int    `json:"credits"`
			Code    string `json:"code"`
		}

		type ClassData struct {
			ID       uint          `json:"id"`
			Name     string        `json:"name"`
			Students []StudentData `json:"students"`
			Courses  []CourseData  `json:"courses"`
		}

		var classes []Model.Class
		var studentData []StudentData
		var courseData []CourseData
		var classData []ClassData

		if err := db.Find(&classes).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "class not found",
				"success": false,
			})
			return
		}

		for _, class := range classes {
			studentData = nil
			courseData = nil
			if err := db.Model(class).Association("Students").Find(&studentData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "student not found",
					"success": false,
				})
				return
			}

			if err := db.Model(class).Association("Courses").Find(&courseData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "course not found",
					"success": false,
				})
				return
			}

			classData = append(classData, ClassData{
				ID:       class.ID,
				Name:     class.Name,
				Students: studentData,
				Courses:  courseData,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    classData,
		})
	})

	// Create class
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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}

		type data struct {
			Name     string `json:"name"`
			CourseID uint   `json:"course_id"`
		}

		var input data
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot bind JSON.",
			})
			return
		}

		newClass := Model.Class{
			Name:      input.Name,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		if err := db.Create(&newClass); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot create class.",
			})
			return
		}

		var course Model.Course
		db.Where("id = ?", input.CourseID).First(&course)

		if err := db.Model(&newClass).Association("Courses").Append(&course); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
				"message": "Class cannot append with course",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			// "message": "Class created.",
			"New Class ID": newClass.ID,
		})
	})

	// Edit class
	r.PATCH("class/:id", Middleware.Authorization(), func(c *gin.Context) {

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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		id, isIdExist := c.Params.Get("id")
		if !isIdExist {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is not supplied.",
			})
			return
		}

		classID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is invalid.",
			})
			return
		}

		var input struct {
			Name     string `json:"name"`
			CourseID uint   `json:"course_id"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot bind JSON.",
			})
			return
		}

		var oldCourse Model.Course
		if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN course_classes ON courses.id = course_classes.course_id").Where("course_classes.class_id = ?", classID).First(&oldCourse).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to fetch course",
			})
			return
		}

		if err := db.Exec("UPDATE course_classes SET course_id = ? WHERE class_id = ?", input.CourseID, classID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to update course_classes association",
			})
			return
		}

		var class Model.Class
		db.Preload("Students").First(&class, "id = ?", classID)

		var newCourse Model.Course
		db.First(&newCourse, "id = ?", input.CourseID)

		for _, student := range class.Students {
			if err := db.Exec("UPDATE course_student SET course_id = ? WHERE course_id = ? AND student_id = ?", input.CourseID, oldCourse.ID, student.ID).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err.Error(),
					"message": "failed to update student_course",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully edited.",
		})
	})

	// Delete class
	r.DELETE("/class/:id", Middleware.Authorization(), func(c *gin.Context) {
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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is not supplied.",
			})
			return
		}
		parsedId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is invalid.",
			})
			return
		}

		delClass := Model.Class{
			ID: uint(parsedId),
		}

		if err := db.First(&delClass, parsedId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Class not found.",
			})
			return
		}

		var delTempCourse Model.Course
		if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN course_classes ON courses.id = course_classes.course_id").Where("course_classes.class_id = ?", delClass.ID).First(&delTempCourse).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch course"})
			return
		}

		if err := db.Model(&delClass).Association("Students").Clear(); err != nil {
			return
		}
		if err := db.Model(&delClass).Association("Courses").Clear(); err != nil {
			return
		}

		if err := db.Model(&delTempCourse).Association("Students").Clear(); err != nil {
			return
		}

		if result := db.Delete(&delClass); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when deleting from the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"message": "successfully deleted.",
		})
	})

	// Add student to class
	r.POST("/addStudentToClass", Middleware.Authorization(), func(c *gin.Context) {
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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
			})
			return
		}

		type inputJSON struct {
			ClassID   uint `json:"class_id"`
			StudentID uint `json:"student_id"`
		}

		var input inputJSON

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot bind JSON class.",
				"error":   err.Error(),
			})
		}

		var class Model.Class
		if err := db.Where("id = ?", input.ClassID).First(&class).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Class not found",
				"error":   err.Error(),
			})
			return
		}

		var student Model.Student
		if err := db.Where("id = ?", input.StudentID).First(&student).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Student not found",
				"error":   err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", input.StudentID).Preload("Courses").First(&student).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Student's courses not found",
				"error":   err.Error(),
			})
			return
		}

		totalCredit := 0

		for _, course := range student.Courses {
			totalCredit += course.Credits
		}

		var studentExists bool
		query := `SELECT EXISTS(SELECT 1 FROM student_classes WHERE student_id = ? AND class_id = ?)`
		if err := db.Raw(query, input.StudentID, input.ClassID).Row().Scan(&studentExists); err != nil {
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

		var course Model.Course
		if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN course_classes ON courses.id = course_classes.course_id").Where("course_classes.class_id = ?", input.ClassID).First(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Courses not found",
				"error":   err.Error(),
			})
			return
		}

		var studentCourse bool
		query = `SELECT EXISTS(SELECT 1 FROM course_student WHERE student_id = ? AND course_id = ?)`
		if err := db.Raw(query, input.StudentID, course.ID).Row().Scan(&studentCourse); err != nil {
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
				"message": "Cannot append class with students",
				"error":   err.Error(),
			})
			return
		}

		if err := db.Model(&course).Association("Students").Append(&student); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot append course with students",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"message": "successfully added new user",
		})
	})

	// Remove student from class
	r.POST("/removeStudentFromClass", Middleware.Authorization(), func(c *gin.Context) {
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

		if user.Role != 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}

		type inputJSON struct {
			ClassID   uint `json:"class_id"`
			StudentID uint `json:"student_id"`
		}

		var input inputJSON

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot bind JSON.",
			})
			return
		}

		if err := db.Exec("DELETE FROM student_classes WHERE student_id = ? AND class_id = ?", input.StudentID, input.ClassID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot delete relation between student and class",
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

		if err := db.Exec("DELETE FROM course_student WHERE student_id = ? AND course_id = ?", input.StudentID, delCourse.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Cannot delete relation between student and course",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			// "success": true,
			"message": "successfully deleted an user",
		})
	})
}
