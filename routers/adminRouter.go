package routers

import (
	"course-management/handlers"
	"course-management/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) *gin.RouterGroup{
	adminGroup := router.Group("/admin")
	{
		adminGroup.POST("/login", handlers.AdminLogin)
		adminGroup.GET("", middleware.AdminRole, middleware.RequireAuth, handlers.AdminGetAllClasses)
		adminGroup.POST("/create-course", middleware.AdminRole, middleware.RequireAuth, handlers.CreateCourse)
		adminGroup.POST("/create-class", middleware.AdminRole, middleware.RequireAuth, handlers.CreateClass)
		adminGroup.PATCH("/class-update", middleware.AdminRole, middleware.RequireAuth, handlers.UpdateClass)
		adminGroup.DELETE("", middleware.AdminRole, middleware.RequireAuth, handlers.DeleteClass)
		adminGroup.GET("/participants", middleware.AdminRole, middleware.RequireAuth, handlers.GetClassParticipants)
		adminGroup.POST("/participants", middleware.AdminRole, middleware.RequireAuth, handlers.AddUserToClass)
		adminGroup.DELETE("/participants",middleware.AdminRole, middleware.RequireAuth, handlers.DeleteUserFromClass)
	}
	return adminGroup
}