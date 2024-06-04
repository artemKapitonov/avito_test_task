package v1

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=operation_history.go -destination=mocks/operation_hisory_mock.go

// OperationHistory defines the interface for getting operation history.
type OperationHistory interface {
	Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error)
}

// getHistory handles the GET request for retrieving operation history.
func (c *Controller) getHistory(ctx *gin.Context) {
	paramID := ctx.Param("id")

	userID, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid 'id' param", err)
		return
	}

	sort, err := selectSortParam(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid 'sort' param", nil)
	}

	isDesc, err := selectDescParam(ctx, sort)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid 'is_descreasing' param", err)
	}

	currency, err := selectCurrency(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "Can't define currency", err)
		return
	}

	operations, err := c.OperationHistory.Get(ctx, userID, sort, isDesc)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, "Can't get user history", err)
		return
	}

	if currency == UsdCurrency {
		fromCurrency := RubCurrency

		for i := range operations {
			operations[i].Amount, err = c.CurrencyConverter.Convert(operations[i].Amount, fromCurrency)
			if err != nil {
				errorResponse(ctx, http.StatusInternalServerError, "Can't do convertation", err)
				return
			}

			operations[i].Currency = currency
		}
	}

	ctx.JSON(http.StatusOK, operations)
}

// selectSortParam is select type of hisory sorting.
func selectSortParam(ctx *gin.Context) (string, error) {
	sort := ctx.Query("sort")
	if sort == "" {
		sort = "date"
	}

	if sort != "amount" && sort != "date" {
		return sort, errors.New("sort param can't be only 'amount' or 'date'")
	}

	return sort, nil
}

// selectDescParam is select decs type of hisory sorting.
func selectDescParam(ctx *gin.Context, sort string) (bool, error) {
	if sort == "amount" {
		isDescParam := ctx.Query("is_descreasing")
		if isDescParam == "" {
			isDescParam = "true"
		}

		return strconv.ParseBool(isDescParam)
	}

	return false, nil
}
