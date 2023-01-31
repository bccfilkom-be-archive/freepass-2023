package server

import (
	"log"

	v1c "github.com/bagashiz/freepass-2023/controllers/api/v1"
	"github.com/bagashiz/freepass-2023/middlewares"
	"github.com/bagashiz/freepass-2023/middlewares/token"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// SetupRouter performs all route operations.
func SetupRouter() {
	token.SetupTokenAuthMiddleware()
	router = gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/students/register", v1c.RegisterStudent)
		v1.POST("/students/login", v1c.LoginStudent)

		v1.POST("/admins/register", v1c.RegisterAdmin)
		v1.POST("/admins/login", v1c.LoginAdmin)
	}

	authV1 := v1.Use(middlewares.AuthMiddleware(token.GetTokenMaker()))
	{
		authV1.GET("/students/:student_id", v1c.GetStudent)
		authV1.PATCH("/students/:student_id", v1c.UpdateStudent)
		authV1.DELETE("/students/:student_id", v1c.DeleteStudent)

		authV1.GET("/admins/:admin_id", v1c.GetAdmin)
		authV1.PATCH("/admins/:admin_id", v1c.UpdateAdmin)
		authV1.DELETE("/admins/:admin_id", v1c.DeleteAdmin)
		authV1.POST("admins/courses/:course_code/classes/:class_code/enroll", v1c.EnrollStudentToClass)
		authV1.DELETE("admins/courses/:course_code/classes/:class_code/enroll", v1c.UnEnrollStudentFromClass)

		authV1.POST("/courses/register", v1c.RegisterCourse)
		authV1.GET("/courses/:course_code", v1c.GetCourse)
		authV1.GET("/courses", v1c.ListCourses)
		authV1.PATCH("/courses/:course_code", v1c.UpdateCourse)
		authV1.DELETE("/courses/:course_code", v1c.DeleteCourse)

		authV1.POST("/courses/:course_code/classes/register", v1c.RegisterClass)
		authV1.GET("/courses/:course_code/classes/:class_code", v1c.GetClass)
		authV1.GET("/courses/:course_code/classes", v1c.ListClasses)
		authV1.PATCH("/courses/:course_code/classes/:class_code", v1c.UpdateClass)
		authV1.DELETE("/courses/:course_code/classes/:class_code", v1c.DeleteClass)

		authV1.POST("/courses/:course_code/classes/:class_code/enroll", v1c.EnrollClass)
		authV1.DELETE("/courses/:course_code/classes/:class_code/enroll", v1c.UnEnrollClass)
	}
}

// Start attaches the router to a server and starts listening and serving HTTP requests from specified address.
func Start(address string) {
	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
