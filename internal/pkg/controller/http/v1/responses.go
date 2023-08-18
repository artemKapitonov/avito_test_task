package v1

import (
	"github.com/gin-gonic/gin"
)

type WrongResponse struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, WrongResponse{msg})
}

type StatusResponse struct {
	Status string `json:"status"`
}
