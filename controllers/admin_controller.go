package controllers

import (
	"bcc_university/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "you did it!!"})
}

func AddUserToClass(c *gin.Context) {
	var payload = c.Param("classId")
	class, err := GetClass(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	var userPayload = c.Param("userId")
	user := GetUser(userPayload)

	_, ok := user.Classes_Enrolled[class.ID.Hex()]
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "user already participated in this class"})
		return
	}

	class.Participants[userPayload] = true
	user.Classes_Enrolled[class.ID.Hex()] = class.Sks
	user.Rem_sks -= class.Sks
	if user.Rem_sks < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "cannot add more class"})
		return
	}

	err2 := mgm.Coll(&models.User{}).Update(user)
	err3 := mgm.Coll(&models.Class{}).Update(class)
	if err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err2.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "successfully added user to class"})
}

func GetUser(userId string) *models.User {
	userPayload2, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	coll := mgm.Coll(user)
	result := coll.FindByID(userPayload2, user)
	if result != nil {
		panic(err)
	}
	return user
}

func DeleteUserFromClass(c *gin.Context) {
	var payload = c.Param("classId")
	class, err := GetClass(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	userId := c.Param("userId")
	user := GetUser(userId)

	delete(class.Participants, userId)
	delete(user.Classes_Enrolled, class.ID.Hex())
	user.Rem_sks += class.Sks

	err2 := mgm.Coll(&models.User{}).Update(user)
	err3 := mgm.Coll(&models.Class{}).Update(class)
	if err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err2.Error()})
		return
	}

	c.JSON(202, gin.H{"message": "successfully deleted user from class"})
}

func CreateClass(c *gin.Context) {
	var payload *models.Class
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	class := &models.Class{}
	coll := mgm.Coll(class)
	createdClass := models.NewClass(payload.Title, payload.Sks)
	err := coll.Create(createdClass)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "successfully created class"})
}

func DeleteClass(c *gin.Context) {
	var payload = c.Param("classId")
	class, err := GetClass(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	for k := range class.Participants {
		user := GetUser(k)
		delete(user.Classes_Enrolled, class.ID.Hex())
		user.Rem_sks += class.Sks
		err = mgm.Coll(&models.User{}).Update(user)

		if err != nil {
			panic(err)
		}
	}
	err = mgm.Coll(&models.Class{}).Delete(class)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(202, gin.H{"message": "successfully deleted class"})
}

func EditClass(c *gin.Context) {
	var payload *EditPayload
	var classId = c.Param("classId")
	class, err := GetClass(classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	edit_map := payload.Edit_Map
	for k, v := range edit_map {
		if k == "title" {
			class.Title = v
		} else if k == "sks" {
			sks, _ := strconv.ParseInt(v, 10, 32)
			class.Sks = int32(sks)
		}
	}

	err = mgm.Coll(&models.Class{}).Update(class)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "successfully edited"})
}
