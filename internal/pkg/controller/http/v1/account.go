package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=account.go -destination=mocks/account_mock.go

// Account is an interface for user accounts
type Account interface {
	Create(ctx context.Context) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

// createUser creates a new user
func (c *Controller) createUser(ctx *gin.Context) {
	user, err := c.Account.Create(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't create user: %s", err.Error()))
		return
	}

	user.Currency = "RUB"

	ctx.JSON(http.StatusOK, user)
}

// userByID retrieves a user by ID
func (c *Controller) userByID(ctx *gin.Context) {
	param := ctx.Param("id")

	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := c.Account.GetByID(ctx, userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't get user by id: %s", err.Error()))
		return
	}

	currency, err := selectCurrency(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if currency == "USD" {
		fromCurrency := "RUB"
		user.Balance, err = c.CurrencyConverter.Convert(user.Balance, fromCurrency, currency)
		if err != nil {
			errorResponse(ctx, http.StatusInternalServerError, "Can't convert currency")
			return
		}
	}

	user.Currency = currency

	ctx.JSON(http.StatusOK, user)
}
