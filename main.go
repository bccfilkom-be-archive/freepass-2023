package main

import (
	"freepass/Config"
	"freepass/Handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	// Connect Database
	db := Config.Connect()
	if db != nil {
		println("Nice, DB Connected")
	}

	// Gin Framework
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(
		func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			if c.Request.Method == "OPTIONS" {
				c.Writer.Header().Set("Content-Type", "application/json")
				c.AbortWithStatus(204)
			} else {
				c.Next()
			}
		},
	)

	r.Group("admin")
	Handler.Admin(db, r)

	r.Group("/api")
	Handler.Register(db, r)
	Handler.Login(db, r)
	Handler.User(db, r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
