package handler

import (
	"net/http"
	"strconv"

	"facebook/admin"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	adminService admin.Service
}

func NewAdminHandler(adminService admin.Service) *adminHandler {
	return &adminHandler{adminService}
}

// fungsi admin menamnbah user
func (h *adminHandler) CreateAdmin(c *gin.Context) {
	var adminRequest admin.AdminRequest

	err := c.ShouldBind(&adminRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	admin, err := h.adminService.Create(adminRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": admin,
	})
}

func (h *adminHandler) GetAdmins(c *gin.Context) {
	admins, err := h.adminService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"data": admins,
	})
}

func (h *adminHandler) GetAdmin(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	admin, err := h.adminService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": admin,
	})
}

func (h *adminHandler) UpdateAdmin(c *gin.Context) {
	var adminRequest admin.AdminRequest

	err := c.ShouldBind(&adminRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	admin, err := h.adminService.Update(id, adminRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": admin,
	})
}

func (h *adminHandler) DeleteAdmin(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	admin, err := h.adminService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": admin,
	})
}
