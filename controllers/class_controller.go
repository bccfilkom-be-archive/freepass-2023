package controllers

import (
	"bcc_university/models"
	"log"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditPayload struct {
	Edit_Map map[string]string `json:"edit_map"`
}

func GetAllClasses(c *gin.Context) {
	class := &models.Class{}
	coll := mgm.Coll(class)
	cur, err := coll.Find(mgm.Ctx(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var results []models.Class
	for cur.Next(mgm.Ctx()) {
		var elem models.Class
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(mgm.Ctx())

	c.JSON(http.StatusOK, gin.H{"classes": results})
}

func AddClass(c *gin.Context) {
	var payload = c.Param("classId")
	currentUser := c.MustGet("currentUser").(*models.User)

	class := &models.Class{}
	coll := mgm.Coll(class)

	payload2, err := primitive.ObjectIDFromHex(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	result := coll.FindByID(payload2, class)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The data doesn't exist"})
		return
	}

	_, ok := currentUser.Classes_Enrolled[class.ID.Hex()]
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "you already participated in this class"})
		return
	}

	currentUser.Classes_Enrolled[class.ID.Hex()] = class.Sks
	currentUser.Rem_sks -= class.Sks
	class.Participants[currentUser.ID.Hex()] = true

	if currentUser.Rem_sks < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "cannot add more class"})
		return
	}
	
	err2 := mgm.Coll(&models.User{}).Update(currentUser)
	err3 := coll.Update(class)
	if err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err2.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "successfully added new class"})
}

func DropClass(c *gin.Context) {
	var payload = c.Param("classId")
	currentUser := c.MustGet("currentUser").(*models.User)

	class := &models.Class{}
	coll := mgm.Coll(class)

	payload2, err := primitive.ObjectIDFromHex(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	result := coll.FindByID(payload2, class)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The data doesn't exist"})
		return
	}

	delete(currentUser.Classes_Enrolled, payload)
	delete(class.Participants, currentUser.ID.Hex())
	currentUser.Rem_sks += class.Sks

	err2 := mgm.Coll(&models.User{}).Update(currentUser)
	err3 := coll.Update(class)
	if err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err2.Error()})
		return
	}
	c.JSON(202, gin.H{"message": "successfully dropped a class"})
}

func ViewParticipants(c *gin.Context) {
	var payload = c.Param("classId")
	class, err := GetClass(payload)
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	user := &models.User{}
	coll := mgm.Coll(user)

	var participants []*models.User
	for k := range class.Participants {
		k2, err := primitive.ObjectIDFromHex(k)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			break
		}
		result2 := coll.FindByID(k2, user)
		if result2 != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The data doesn't exist"})
			break
		} else {
			participants = append(participants, user)
		}
	}

	c.JSON(http.StatusOK, gin.H{"participants": participants})
}

func GetClass(classId string) (*models.Class, error) {
	classId2, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return nil, err
	}

	class := &models.Class{}
	coll := mgm.Coll(class)
	err = coll.FindByID(classId2, class)
	if err != nil {
		return nil, fmt.Errorf("cannot find class")
	}
	return class, nil
}
