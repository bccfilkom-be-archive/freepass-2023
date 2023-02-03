package handler

import (
	"net/http"
	"strconv"

	"facebook/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// fungsi user melakukan register
func (h *userHandler) CreateUser(c *gin.Context) {
	var userRequest user.UserRequest

	err := c.ShouldBind(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	user, err := h.userService.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"data": users,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// fungsi user update dirinya sendiri
func (h *userHandler) UpdateUser(c *gin.Context) {
	var userRequest user.UserRequest

	err := c.ShouldBind(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.Update(id, userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) DeleteAdmin(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
