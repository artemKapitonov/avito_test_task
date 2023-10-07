package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"

	"github.com/gin-gonic/gin"
)

const (
	UsdCurrency = "USD"
	RubCurrency = "RUB"
)

//go:generate mockgen -source=account.go -destination=mocks/account_mock.go

// Account is an interface for user accounts.
type Account interface {
	Create(ctx context.Context) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

// createUser creates a new user.
func (c *Controller) createUser(ctx *gin.Context) {
	user, err := c.Account.Create(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, "Can't create user", err)
		return
	}

	user.Currency = RubCurrency

	ctx.JSON(http.StatusOK, user)
}

// userByID retrieves a user by ID.
func (c *Controller) userByID(ctx *gin.Context) {
	param := ctx.Param("id")

	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param", err)
		return
	}

	user, err := c.Account.GetByID(ctx, userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, "Can't get user by id", err)
		return
	}

	currency, err := selectCurrency(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "Can't define currency", err)
		return
	}

	user.Currency = currency

	if currency == UsdCurrency {
		fromCurrency := RubCurrency

		user.Balance, err = c.CurrencyConverter.Convert(user.Balance, fromCurrency)
		if err != nil {
			errorResponse(ctx, http.StatusInternalServerError, "Can't convert currency", err)
			return
		}
	}

	ctx.JSON(http.StatusOK, user)
}
