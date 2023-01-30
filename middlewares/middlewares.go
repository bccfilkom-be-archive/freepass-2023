package middlewares

import (
	"bcc_university/initializers"
	"bcc_university/models"
	"bcc_university/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, _ := initializers.LoadConfig(".")
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user := &models.User{}
		coll := mgm.Coll(user)
		result := coll.FindByID(fmt.Sprint(sub), user)
		if result != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no longer exists", "sub": sub})
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet("currentUser").(*models.User)
		if currentUser.Groups != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "unauthorized"})
		} else {
			c.Next()
		}
	}
}




