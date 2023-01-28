package routers

import (
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) *gin.RouterGroup{
	adminGroup := router.Group("/admin")
	
	return adminGroup
}