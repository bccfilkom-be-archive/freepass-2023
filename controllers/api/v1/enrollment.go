package v1controllers

import (
	"errors"
	"net/http"

	"github.com/bagashiz/freepass-2023/helpers"
	"github.com/bagashiz/freepass-2023/middlewares"
	"github.com/bagashiz/freepass-2023/middlewares/token"
	v1s "github.com/bagashiz/freepass-2023/services/api/v1"
	"github.com/gin-gonic/gin"
)

// updateEnrollmentRequest is a struct for updating an enrollment.
type updateEnrollmentRequest struct {
	ClassCode  string `uri:"class_code" binding:"required"`
	CourseCode string `uri:"course_code" binding:"required"`
}

// EnrollClass is a function that enrolls a student to a class.
func EnrollClass(ctx *gin.Context) {
	var req updateEnrollmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateEnrollmentParams{
		ClassCode:  req.ClassCode,
		CourseCode: req.CourseCode,
		StudentID:  authPayload.UserID,
	}

	student, err := v1s.EnrollStudentToClass(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsInvalidData(err) {
			err = errors.New("class is full, student already enrolled, student credits exceeds 24")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// UnEnrollClass is a function that unenrolls a student from a class.
func UnEnrollClass(ctx *gin.Context) {
	var req updateEnrollmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateEnrollmentParams{
		ClassCode:  req.ClassCode,
		CourseCode: req.CourseCode,
		StudentID:  authPayload.UserID,
	}

	student, err := v1s.UnEnrollStudentFromClass(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsInvalidData(err) {
			err = errors.New("student has not enrolled")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// studentIdJSON is a struct to contain student id for enrollment by admin.
type studentIdJSON struct {
	StudentID string `json:"student_id" binding:"required"`
}

// EnrollStudentToClass is a function that enrolls a student to a class by admin.
func EnrollStudentToClass(ctx *gin.Context) {
	var uri updateEnrollmentRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	var json studentIdJSON
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateEnrollmentParams{
		ClassCode:  uri.ClassCode,
		CourseCode: uri.CourseCode,
		StudentID:  json.StudentID,
	}

	student, err := v1s.EnrollStudentToClass(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsInvalidData(err) {
			err = errors.New("class is full, student already enrolled, student credits exceeds 24")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// UnEnrollStudentFromClass is a function that unenrolls a student from a class by admin.
func UnEnrollStudentFromClass(ctx *gin.Context) {
	var uri updateEnrollmentRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	var json studentIdJSON
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateEnrollmentParams{
		ClassCode:  uri.ClassCode,
		CourseCode: uri.CourseCode,
		StudentID:  json.StudentID,
	}

	student, err := v1s.UnEnrollStudentFromClass(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsInvalidData(err) {
			err = errors.New("student has not enrolled")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newStudentResponse(student)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}
