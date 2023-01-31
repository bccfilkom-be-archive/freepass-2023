package v1controllers

import (
	"errors"
	"net/http"

	"github.com/bagashiz/freepass-2023/helpers"
	"github.com/bagashiz/freepass-2023/middlewares"
	"github.com/bagashiz/freepass-2023/middlewares/token"
	"github.com/bagashiz/freepass-2023/models"
	v1s "github.com/bagashiz/freepass-2023/services/api/v1"
	"github.com/gin-gonic/gin"
)

// classResponse is a struct for returning a class.
type classResponse struct {
	ClassCode  string           `json:"class_code"`
	CourseCode string           `json:"course_code"`
	Capacity   int              `json:"capacity"`
	Enrolled   int              `json:"enrolled"`
	Students   []studentInClass `json:"students"`
}

// studentInClass is a struct for returning a student in a class.
type studentInClass struct {
	StudentID string `json:"student_id"`
	FullName  string `json:"full_name"`
}

// newClassResponse is a function that returns a classResponse.
func newClassResponse(class models.Class) *classResponse {
	var students []studentInClass
	for _, student := range class.Students {
		students = append(students, studentInClass{
			StudentID: student.StudentID,
			FullName:  student.FullName,
		})
	}

	return &classResponse{
		ClassCode:  class.ClassCode,
		CourseCode: class.CourseCode,
		Capacity:   class.Capacity,
		Enrolled:   class.Enrolled,
		Students:   students,
	}
}

// registerClassRequest is a struct for creating a new class.
type registerClassRequest struct {
	ClassCode string `json:"class_code" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required,number,min=1,max=40"`
}

// RegisterClass is a function that creates a new class.
func RegisterClass(ctx *gin.Context) {
	var req registerClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	courseCode := ctx.Param("course_code")

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.RegisterClassParams{
		ClassCode:  req.ClassCode,
		CourseCode: courseCode,
		Capacity:   req.Capacity,
	}

	class, err := v1s.RegisterClass(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			err = errors.New("course not found")
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsUniqueViolation(err) {
			err = errors.New("class already registered")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsValueTooLong(err) {
			err = errors.New("class code should only contain 1 character")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newClassResponse(class)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// getClassRequest is a struct for getting a class.
type getClassRequest struct {
	ClassCode  string `uri:"class_code" binding:"required"`
	CourseCode string `uri:"course_code" binding:"required"`
}

// GetClass is a function that returns a class based on class code.
func GetClass(ctx *gin.Context) {
	var req getClassRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	class, err := v1s.GetClass(ctx, req.ClassCode)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	// if the user is not enrolled and not an admin, set class.Students to nil
	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	var isEnrolled bool
	for _, student := range class.Students {
		if authPayload.UserID == student.StudentID {
			isEnrolled = true
			break
		}
	}

	if !isEnrolled && !authPayload.IsAdmin {
		class.Students = nil
	}

	rsp := newClassResponse(class)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// listClassesRequest is a struct for listing classes with pagination.
type listClassesRequest struct {
	Page int `form:"page" binding:"required,number,min=1"`
	Size int `form:"size" binding:"required,number,min=1,max=20"`
}

// ListClasses is a function that returns a list of classes with pagination and sorted by code.
func ListClasses(ctx *gin.Context) {
	var req listClassesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.ListClasesParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	classes, err := v1s.ListClasses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	// if the user is not enrolled and not an admin, set class.Students to nil
	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	for i, class := range classes {
		var isEnrolled bool
		for _, student := range class.Students {
			if authPayload.UserID == student.StudentID {
				isEnrolled = true
				break
			}
		}

		if !isEnrolled && !authPayload.IsAdmin {
			classes[i].Students = nil
		}
	}

	rsp := make([]*classResponse, len(classes))
	for i, class := range classes {
		rsp[i] = newClassResponse(class)
	}

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// updateClassRequest is a struct for updating a class.
type updateClassRequest struct {
	Capacity int `json:"capacity" binding:"required,number,min=1,max=40"`
}

// UpdateClass is a function that updates a class based on class code.
func UpdateClass(ctx *gin.Context) {
	var req updateClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	classCode := ctx.Param("class_code")
	courseCode := ctx.Param("course_code")

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.UpdateClassParams{
		ClassCode:  classCode,
		CourseCode: courseCode,
		Capacity:   req.Capacity,
	}

	class, err := v1s.UpdateClass(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsInvalidData(err) {
			err = errors.New("new capacity must be greater than current enrolled")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newClassResponse(class)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// deleteClassRequest is a struct for deleting a class.
type deleteClassRequest struct {
	ClassCode  string `uri:"class_code" binding:"required"`
	CourseCode string `uri:"course_code" binding:"required"`
}

// DeleteClass is a function that deletes a class based on class code.
func DeleteClass(ctx *gin.Context) {
	var req deleteClassRequest
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

	class, err := v1s.DeleteClass(ctx, req.ClassCode)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newClassResponse(class)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}
