package main

import (
	"bcc_university/controllers"
	"bcc_university/middlewares"

	"bcc_university/initializers"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	config, _ := initializers.LoadConfig(".")
	_ = mgm.SetDefaultConfig(nil, "bcc_university", options.Client().ApplyURI(config.MongoUri))
}

func main() {
	Init()
	router := gin.Default()

	router.POST("/register", controllers.SignUpUser)
	router.POST("/login", controllers.SignInUser)
	
	adminRoutes := router.Group("/class")
	adminRoutes.Use(middlewares.JwtAuthMiddleware(), middlewares.AdminAuth())
	{
		adminRoutes.POST("/:classId/add-user/:userId", controllers.AddUserToClass)
		adminRoutes.DELETE("/:classId/delete-user/:userId", controllers.DeleteUserFromClass)
		adminRoutes.POST("/create", controllers.CreateClass)
		adminRoutes.PUT("/:classId/edit", controllers.EditClass)
		adminRoutes.DELETE("/delete/:classId", controllers.DeleteClass)
	}
	
	authorized := router.Group("/")
	authorized.Use(middlewares.JwtAuthMiddleware())
	{
		authorized.PUT("/user/profile/edit", controllers.EditProfile)
		authorized.GET("/logout", controllers.LogoutUser)
		authorized.GET("/user/profile", controllers.GetUserDetails)
		authorized.GET("/class", controllers.GetAllClasses)
		authorized.POST("myclass/add-class/:classId", controllers.AddClass)
		authorized.GET("/myclass/:classId/participants", controllers.ViewParticipants)
		authorized.DELETE("/myclass/delete-class/:classId", controllers.DropClass)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// docker compose --env-file app.env up