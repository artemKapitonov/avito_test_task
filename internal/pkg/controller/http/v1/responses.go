// Package v1 provides functionality for handling HTTP requests and responses.
package v1

import (
	"github.com/gin-gonic/gin"
)

// WrongResponse is a struct that represents an error response.
type WrongResponse struct {
	Error string `json:"error" example:"message"`
}

// errorResponse is a function that sends an error response with the given message and HTTP status code.
func errorResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, WrongResponse{msg})
}

// StatusResponse is a struct that represents a status response.
type StatusResponse struct {
	Status string `json:"status"`
}
