// Package v1 provides functionality for handling HTTP requests and responses.
package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// WrongResponse is a struct that represents an error response.
type WrongResponse struct {
	Error string `json:"error,omitempty"`
}

// errorResponse is a function that sends an error response with the given message and HTTP status code.
func errorResponse(ctx *gin.Context, code int, msg string, err error) {
	if err != nil {
		ctx.AbortWithStatusJSON(code, WrongResponse{fmt.Sprintf(msg, "Error:", err.Error())})
	}

	ctx.AbortWithStatusJSON(code, WrongResponse{msg})
}

// StatusResponse is a struct that represents a status response.
type StatusResponse struct {
	Status string `json:"status"`
}
