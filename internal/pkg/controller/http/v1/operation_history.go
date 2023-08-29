package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=operation_history.go -destination=mock/operation_hisory_mock.go

type OperationHistory interface {
	Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error)
}

func (c *Controller) getHistory(ctx *gin.Context) {
	paramID := ctx.Param("id")

	userID, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	sort := ctx.Query("sort")

	if sort == "" {
		sort = "date"
	}

	if sort != "amount" && sort != "date" {
		errorResponse(ctx, http.StatusBadRequest, "invalid 'sort' param")
		return
	}

	var isDesc bool

	if sort == "amount" {
		isDescParam := ctx.Query("is_descreasing")

		if isDescParam == "" {
			isDescParam = "true"
		}

		isDesc, err = strconv.ParseBool(isDescParam)
		if err != nil {
			errorResponse(ctx, http.StatusBadRequest, "invalid is_descreasing param")
		}
	}

	currency, err := selectCurrency(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	operations, err := c.OperationHistory.Get(ctx, userID, sort, isDesc)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if currency == "USD" {
		fromCurrency := "RUB"

		for i := range operations {
			operations[i].Amount, err = c.CurrencyConverter.Convert(operations[i].Amount, fromCurrency, currency)
			if err != nil {
				errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't do convertation: %s", err.Error()))
			}
			operations[i].Currency = currency
		}
	}

	ctx.JSON(http.StatusOK, operations)
}
