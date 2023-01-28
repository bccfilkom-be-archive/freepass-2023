package routers

import (
	"github.com/gin-gonic/gin"
	"course-management/handlers"
	"course-management/middleware"
)

func UserRouter(router *gin.Engine) *gin.RouterGroup{
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", handlers.UserSignup)
		userGroup.POST("/login", handlers.UsersLogin)
		userGroup.GET("/profile", middleware.UserRole, middleware.RequireAuth, handlers.GetUserProfile)
		userGroup.PATCH("/profile", middleware.UserRole, middleware.RequireAuth, handlers.UpdateProfile)
		userGroup.GET("", middleware.UserRole, middleware.RequireAuth, handlers.GetAllClasses)
		userGroup.GET("/search", middleware.UserRole, middleware.RequireAuth, handlers.SearchClass)
		userGroup.POST("", middleware.UserRole, middleware.RequireAuth, handlers.AddClass)
		userGroup.GET("/class", middleware.UserRole, middleware.RequireAuth, handlers.GetUserClasses)
		userGroup.GET("/class/participants", middleware.UserRole, middleware.RequireAuth, handlers.GetClassParticipants)
		userGroup.DELETE("class/:id",middleware.UserRole, middleware.RequireAuth, handlers.DeleteUserClass)
	}
	return userGroup
}