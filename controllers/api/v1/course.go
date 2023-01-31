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

// courseResponse is a struct for returning a course.
type courseResponse struct {
	CourseCode string          `json:"course_code"`
	CourseName string          `json:"course_name"`
	Credits    int             `json:"credits"`
	Classes    []classInCourse `json:"classes"`
}

// classInCourse is a struct for returning a class in a course.
type classInCourse struct {
	ClassCode string `json:"class_code"`
	Capacity  int    `json:"capacity"`
	Enrolled  int    `json:"enrolled"`
}

// newCourseResponse is a function that returns a courseResponse.
func newCourseResponse(course models.Course) *courseResponse {
	var classes []classInCourse
	for _, class := range course.Classes {
		classes = append(classes, classInCourse{
			ClassCode: class.ClassCode,
			Capacity:  class.Capacity,
			Enrolled:  class.Enrolled,
		})
	}

	return &courseResponse{
		CourseCode: course.CourseCode,
		CourseName: course.CourseName,
		Credits:    course.Credits,
		Classes:    classes,
	}
}

// registerCourseRequest is a struct for creating a new course.
type registerCourseRequest struct {
	CourseCode string `json:"course_code" binding:"required"`
	CourseName string `json:"course_name" binding:"required"`
	Credits    int    `json:"credits" binding:"required,number,min=1,max=6"`
}

// RegisterCourse is a function that creates a new course.
func RegisterCourse(ctx *gin.Context) {
	var req registerCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(middlewares.GetAuthPayload()).(*token.Payload)
	if !authPayload.IsAdmin {
		err := errors.New("unauthorized access")
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.RegisterCourseParams{
		CourseCode: req.CourseCode,
		CourseName: req.CourseName,
		Credits:    req.Credits,
	}

	course, err := v1s.RegisterCourse(ctx, arg)
	if err != nil {
		if helpers.IsUniqueViolation(err) {
			err = errors.New("course already registered")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsValueTooLong(err) {
			err = errors.New("course code should be less than 8 characters")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newCourseResponse(course)

	ctx.JSON(http.StatusCreated, helpers.Response(rsp))
}

// getCourseRequest is a struct for getting a course.
type getCourseRequest struct {
	CourseCode string `uri:"course_code" binding:"required"`
}

// GetCourse is a function that gets a course.
func GetCourse(ctx *gin.Context) {
	var req getCourseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	course, err := v1s.GetCourse(ctx, req.CourseCode)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newCourseResponse(course)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// listCoursesRequest is a struct for getting a list of courses with pagination.
type listCoursesRequest struct {
	Page int `form:"page" binding:"required,number,min=1"`
	Size int `form:"size" binding:"required,number,min=5,max=10"`
}

// ListCourses is a function that gets a list of courses with pagination.
func ListCourses(ctx *gin.Context) {
	var req listCoursesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	arg := v1s.ListCoursesParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	courses, err := v1s.ListCourses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := make([]*courseResponse, len(courses))
	for i, course := range courses {
		rsp[i] = newCourseResponse(course)
	}

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// updateCourseRequest is a struct for updating a course.
type updateCourseRequest struct {
	CourseName string `json:"course_name" binding:"required"`
	Credits    int    `json:"credits" binding:"required,number,min=1,max=6"`
}

// UpdateCourse is a function that updates a course.
func UpdateCourse(ctx *gin.Context) {
	var req updateCourseRequest
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

	arg := v1s.UpdateCourseParams{
		CourseCode: courseCode,
		CourseName: req.CourseName,
		Credits:    req.Credits,
	}

	course, err := v1s.UpdateCourse(ctx, arg)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsInvalidData(err) {
			err = errors.New("course already has class(es), cannot update")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newCourseResponse(course)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}

// deleteCourseRequest is a struct for deleting a course.
type deleteCourseRequest struct {
	CourseCode string `uri:"course_code" binding:"required"`
}

// DeleteCourse is a function that deletes a course.
func DeleteCourse(ctx *gin.Context) {
	var req deleteCourseRequest
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

	course, err := v1s.DeleteCourse(ctx, req.CourseCode)
	if err != nil {
		if helpers.IsRecordNotFound(err) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}

		if helpers.IsForeignKeyViolation(err) {
			err = errors.New("course already assigned to class(es), delete the class(es) first")
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := newCourseResponse(course)

	ctx.JSON(http.StatusOK, helpers.Response(rsp))
}
