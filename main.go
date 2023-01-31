package main

import (
	"course-management/initializers"
	"course-management/routers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.Syncdatabase()
}

func main(){
	router :=gin.Default()
	
	routers.UserRouter(router)
	routers.AdminRouter(router)

	router.Run()
}
