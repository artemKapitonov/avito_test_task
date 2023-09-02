package v1

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=currency.go -destination=mocks/currency_mock.go

// CurrencyConverter is an interface for converting currencies
type CurrencyConverter interface {
	Convert(amount float64, fromCurrency, toCurrency string) (float64, error)
}

// selectCurrency selects the currency based on the query parameter in the context
func selectCurrency(ctx *gin.Context) (string, error) {
	currencyParam := ctx.Query("currency")

	switch strings.ToUpper(currencyParam) {
	case "":
		return "RUB", nil

	case "RUB":
		return "RUB", nil

	case "USD":
		return "USD", nil

	default:
		return "", errors.New("invalid currency parameter")
	}
}
