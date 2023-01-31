package v1controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/bagashiz/freepass-2023/helpers"
	"github.com/bagashiz/freepass-2023/middlewares"
	"github.com/bagashiz/freepass-2023/middlewares/token"
	"github.com/bagashiz/freepass-2023/models"
	v1s "github.com/bagashiz/freepass-2023/services/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// adminResponse is a struct for returning admin data.
type adminResponse struct {
	AdminID   string    `json:"admin_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// newAdminResponse is a function that creates a new adminResponse.
func newAdminResponse(admin models.Admin) adminResponse {
	return adminResponse{
		AdminID:   admin.AdminID,
		FullName:  admin.FullName,
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

// registerAdminRequest is a struct for registering a new admin user.
type registerAdminRequest struct {
	AdminID  string `json:"admin_id" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterAdmin is a function that registers a new admin user.
func RegisterAdmin(ctx *gin.Context) {
	var req registerAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.RegisterAdminParams{
		AdminID:        req.AdminID,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	admin, err := v1s.RegisterAdmin(ctx, arg)
	if err != nil {
		if helpers.IsUniqueViolation(err) {
			err = errors.New("admin_id or email already registered")
			ctx.JSON(http.StatusConflict, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newAdminResponse(admin)

	ctx.JSON(http.StatusCreated, helpers.Response(rsp))
}

// loginAdminRequest is a struct for logging in an admin user.
type loginAdminRequest struct {
	AdminID  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// loginAdminResponse is a struct for returning login data.
type loginAdminResponse struct {
	AccessToken          string        `json:"access_token"`
	AccessTokenExpiresAt time.Time     `json:"access_token_expires_at"`
	Admin                adminResponse `json:"admin"`
}

// LoginAdmin is a function that logs in an admin user.
func LoginAdmin(ctx *gin.Context) {
	var req loginAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	admin, err := v1s.GetAdmin(ctx, req.AdminID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	err = helpers.ComparePassword(req.Password, admin.HashedPassword)
	if err != nil {
		err = errors.New("invalid password")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	isAdmin := true
	accessToken, accessPayload, err := token.GetTokenMaker().CreateToken(
		admin.AdminID,
		token.GetTokenDuration(),
		isAdmin,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := loginAdminResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		Admin:                newAdminResponse(admin),
	}

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// getAdminRequest is a struct for getting an admin user based on Admin ID.
type getAdminRequest struct {
	AdminID string `uri:"admin_id" binding:"required"`
}

// GetAdmin is a function that gets an admin user based on Admin ID.
func GetAdmin(ctx *gin.Context) {
	var req getAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if req.AdminID != authPayload.UserID && !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	admin, err := v1s.GetAdmin(ctx, req.AdminID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newAdminResponse(admin)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// updateAdminRequest is a struct for updating an admin user based on NIP.
type updateAdminRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateAdmin is a function that updates an admin user based on NIP.
func UpdateAdmin(ctx *gin.Context) {
	var req updateAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	adminID := ctx.Param("admin_id")

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if adminID != authPayload.UserID && !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateAdminParams{
		AdminID:        adminID,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	admin, err := v1s.UpdateAdmin(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newAdminResponse(admin)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// deleteAdminRequest is a struct for deleting an admin user based on AdminID.
type deleteAdminRequest struct {
	AdminID string `uri:"admin_id" binding:"required"`
}

// deleteAdminResponse is a struct for deleting an admin user based on AdminID.
type deleteAdminResponse struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	adminResponse
}

// DeleteAdmin is a function that deletes an admin user based on AdminID.
func DeleteAdmin(ctx *gin.Context) {
	var req deleteAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if authPayload.IsAdmin {
		if req.AdminID == authPayload.UserID {
			err := errors.New("cannot delete your own account")
			ctx.JSON(http.StatusForbidden, helpers.ErrorResponse(err))
			return
		}
	} else {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	admin, err := v1s.DeleteAdmin(ctx, req.AdminID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := deleteAdminResponse{
		DeletedAt:     admin.DeletedAt,
		adminResponse: newAdminResponse(admin),
	}

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}
