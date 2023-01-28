package middleware

import (
	"course-management/initializers"
	"course-management/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func UserRole(c *gin.Context){
	c.Set("role","user")
	c.Next()
}

func AdminRole(c *gin.Context){
	c.Set("role","admin")
	c.Next()
}

func RequireAuth(c *gin.Context) {
	tokenString,err:=c.Cookie("Authorization")
	if err !=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRETTOKEN")), nil
	})


	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		role,_:=c.Get("role")
		if claims["role"]!=role{
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		switch claims["role"] {
		case "user":
			var user models.Users
			initializers.DB.First(&user,claims["sub"])
			if user.ID==0{
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Set("user",user.ID)
		case "admin":
			var admin models.Admin
			initializers.DB.First(&admin,claims["sub"])
			if admin.ID==0{
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Set("admin",admin.ID)
		}
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}