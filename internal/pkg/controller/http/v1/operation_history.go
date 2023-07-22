package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/gin-gonic/gin"
)

type OperationHistory interface {
	Get(ctx context.Context, userID uint64) ([]entity.Operation, error)
}

func (c *Controller) getHistory(ctx *gin.Context) {
	paramID := ctx.Param("id")

	userID, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	operations, err := c.OperationHistory.Get(ctx, userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, operations)
}
