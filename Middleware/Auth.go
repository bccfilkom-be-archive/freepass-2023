package Middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = header[len("Bearer "):]
		godotenv.Load("../.env")
		token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_G")), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "JWT validation error.",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", uint(claims["id"].(float64)))
			c.Next()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "JWT invalid.",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
	}
}
