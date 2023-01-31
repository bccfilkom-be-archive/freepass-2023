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

// studentResponse is a struct for returning student data.
type studentResponse struct {
	StudentID    string            `json:"student_id"`
	FullName     string            `json:"full_name"`
	Email        string            `json:"email"`
	TotalCredits int               `json:"total_credits"`
	Classes      []classesEnrolled `json:"classes"`
}

// classesEnrolled is a struct for returning classes enrolled by a student.
type classesEnrolled struct {
	ClassCode  string `json:"class_code"`
	CourseCode string `json:"course_code"`
}

// newStudentResponse is a function that creates a new studentResponse.
func newStudentResponse(student models.Student) studentResponse {
	var classes []classesEnrolled
	for _, class := range student.Classes {
		classes = append(classes, classesEnrolled{
			ClassCode:  class.ClassCode,
			CourseCode: class.CourseCode,
		})
	}

	return studentResponse{
		StudentID:    student.StudentID,
		FullName:     student.FullName,
		Email:        student.Email,
		TotalCredits: student.TotalCredits,
		Classes:      classes,
	}
}

// registerStudentRequest is a struct for registering a new student user.
type registerStudentRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	FullName  string `json:"full_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

// RegisterStudent is a function that registers a new student user.
func RegisterStudent(ctx *gin.Context) {
	var req registerStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.RegisterStudentParams{
		StudentID:      req.StudentID,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	student, err := v1s.RegisterStudent(ctx, arg)
	if err != nil {
		if helpers.IsUniqueViolation(err) {
			err = errors.New("student_id or email already registered")
			ctx.JSON(http.StatusConflict, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusCreated, helpers.Response(rsp))
}

// loginStudentRequest is a struct for logging in a student user.
type loginStudentRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// loginStudentResponse is a struct for login response.
type loginStudentResponse struct {
	AccessToken          string          `json:"access_token"`
	AccessTokenExpiresAt time.Time       `json:"access_token_expires_at"`
	Student              studentResponse `json:"student"`
}

// LoginStudent is a function that logs in a student user.
func LoginStudent(ctx *gin.Context) {
	var req loginStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	student, err := v1s.GetStudent(ctx, req.StudentID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	err = helpers.ComparePassword(req.Password, student.HashedPassword)
	if err != nil {
		err = errors.New("invalid password")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	isAdmin := false
	accessToken, accessPayload, err := token.GetTokenMaker().CreateToken(
		student.StudentID,
		token.GetTokenDuration(),
		isAdmin,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := loginStudentResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		Student:              newStudentResponse(student),
	}

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// getStudentRequest is a struct for getting a student user based on StudentID.
type getStudentRequest struct {
	StudentID string `uri:"student_id" binding:"required"`
}

// GetStudent is a function that returns a student user based on StudentID.
func GetStudent(ctx *gin.Context) {
	var req getStudentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if req.StudentID != authPayload.UserID && !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	student, err := v1s.GetStudent(ctx, req.StudentID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// updateStudentRequest is a struct for updating a student user.
type updateStudentRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

// UpdateStudent is a function that updates a student user's email and password.
func UpdateStudent(ctx *gin.Context) {
	var req updateStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	studentID := ctx.Param("student_id")

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if studentID != authPayload.UserID {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateStudentParams{
		StudentID:      studentID,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	student, err := v1s.UpdateStudent(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// deleteStudentRequest is a struct for deleting a student user based on StudentID.
type deleteStudentRequest struct {
	StudentID string `uri:"student_id" binding:"required"`
}

// deleteStudentResponse is a struct for deleting a student user.
type deleteStudentResponse struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	studentResponse
}

// DeleteStudent is a function that deletes a student user based on StudentID.
func DeleteStudent(ctx *gin.Context) {
	var req deleteStudentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	student, err := v1s.DeleteStudent(ctx, req.StudentID)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := deleteStudentResponse{
		DeletedAt:       student.DeletedAt,
		studentResponse: newStudentResponse(student),
	}

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}
