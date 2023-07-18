package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	operationTransfer = "transfer"
	operationAccrual  = "accrual"
	operationDebit    = "redeem"
)

type Balance interface {
	Update(ctx context.Context, userID uint64, amount float64) error
	Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error
}

func (c *Controller) updateBalance(ctx *gin.Context) {
	param := ctx.Param("id")

	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	query := ctx.Query("amount")
	amount, err := strconv.ParseFloat(query, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.Balance.Update(ctx, userID, amount); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't do operation: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "succes",
	})
}

func (c *Controller) transfer(ctx *gin.Context) {
	param := ctx.Param("sender_id")

	senderID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	param = ctx.Param("recipient_id")

	recipientID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	query := ctx.Query("amount")
	amount, err := strconv.ParseFloat(query, 64)
	if err != nil || amount <= 0 {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.Balance.Transfer(ctx, senderID, recipientID, amount); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't make a transfer: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "succes",
	})
}
