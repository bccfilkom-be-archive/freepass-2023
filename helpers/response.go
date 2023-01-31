package helpers

import (
	"github.com/gin-gonic/gin"
)

// Response is a function that returns common format for API responses.
func Response(data interface{}) gin.H {
	return gin.H{"data": data}
}

// ErrorResponse is a function that returns common format for API error responses.
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
