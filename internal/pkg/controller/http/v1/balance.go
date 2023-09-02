package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=balance.go -destination=mocks/balance_mock.go

// Balance represents the balance interface
type Balance interface {
	Update(ctx context.Context, userID uint64, amount float64) error
	Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error
}

// updateBalance updates the balance of a user
func (c *Controller) updateBalance(ctx *gin.Context) {
	// Get the user ID from the URL parameter
	param := ctx.Param("id")

	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	// Get the amount from the query parameter
	query := ctx.Query("amount")
	amount, err := strconv.ParseFloat(query, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid amount param")
		return
	}

	// Call the Update method of the Balance interface
	if err := c.Balance.Update(ctx, userID, amount); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't do operation: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "success",
	})
}

// transfer transfers an amount from one user to another
func (c *Controller) transfer(ctx *gin.Context) {
	// Get the sender ID from the URL parameter
	param := ctx.Param("sender_id")

	senderID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid sender id param")
		return
	}

	// Get the recipient ID from the URL parameter
	param = ctx.Param("recipient_id")

	recipientID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid recipient id param")
		return
	}

	// Get the amount from the query parameter
	query := ctx.Query("amount")
	amount, err := strconv.ParseFloat(query, 64)
	if err != nil || amount <= 0 {
		errorResponse(ctx, http.StatusBadRequest, "invalid amount param")
		return
	}

	// Get the currency from the request
	fromCurrency, err := selectCurrency(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Convert the amount if the currency is USD
	if fromCurrency == "USD" {
		toCurrency := "RUB"
		amount, err = c.CurrencyConverter.Convert(amount, fromCurrency, toCurrency)
		if err != nil {
			errorResponse(ctx, http.StatusInternalServerError, "Can't convert currency")
			return
		}
	}

	// Call the Transfer method of the Balance interface
	if err := c.Balance.Transfer(ctx, senderID, recipientID, amount); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't make a transfer: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "success",
	})
}
