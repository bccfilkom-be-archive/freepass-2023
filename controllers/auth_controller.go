package controllers

import (
	"bcc_university/initializers"
	"bcc_university/models"
	"bcc_university/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"

	"go.mongodb.org/mongo-driver/bson"
)

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func SignUpUser(c *gin.Context) {
	var payload *models.User
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user := &models.User{}
	coll := mgm.Coll(user)

	takenUsername := coll.First(bson.M{"username": payload.Username}, user)
	takenEmail := coll.First(bson.M{"email": payload.Email}, user)

	if takenUsername == nil || takenEmail == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username or email has already taken"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	createdUser := models.NewUser(payload.Username, payload.Email, hashedPassword)
	coll.Create(createdUser)
	c.JSON(http.StatusCreated, gin.H{"username": createdUser.Username})
}

func SignInUser(c *gin.Context) {
	var payload *SignInInput

	user := &models.User{}
	coll := mgm.Coll(user)

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result := coll.First(bson.M{"email": strings.ToLower(payload.Email)}, user)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")
	tokenExpires := 1 * time.Hour

	// Generate Token
	token, err2 := utils.GenerateToken(tokenExpires, user.ID, config.TokenSecret)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err2.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}

func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Logout success"})
}

func GetUserDetails(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*models.User)
	userDetails := &models.UserDetails{
		Username:        currentUser.Username,
		Email:           currentUser.Email,
		FirstName:       currentUser.FirstName,
		LastName:        currentUser.LastName,
		ClassesEnrolled: currentUser.ClassesEnrolled,
	}

	c.JSON(http.StatusOK, userDetails)
}

func EditProfile(c *gin.Context) {
	var payload *EditPayload
	currentUser := c.MustGet("currentUser").(*models.User)
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	editMap := payload.EditMap
	for k, v := range editMap {
		if k == "first_name" {
			currentUser.FirstName = v
		} else if k == "last_name" {
			currentUser.LastName = v
		} else if k == "email" {
			currentUser.Email = v
		} else if k == "username" {
			currentUser.Username = v
		}
	}

	err := mgm.Coll(&models.User{}).Update(currentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "successfully edited"})
}
