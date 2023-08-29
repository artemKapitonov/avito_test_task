package v1

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=currency.go -destination=mock/currency_mock.go

type CurrencyConverter interface {
	Convert(amount float64, fromCurrency, toCurrency string) (float64, error)
}

func selectCurrency(ctx *gin.Context) (string, error) {
	currencyParam := ctx.Query("currency")

	switch strings.ToTitle(currencyParam) {
	case "":
		return "RUB", nil

	case "RUB":
		return "RUB", nil

	case "USD":
		return "USD", nil

	default:
		return "", errors.New("invalid currency param")
	}
}
