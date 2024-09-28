package currencyconverter

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

var (
	rubToUsdRate float64
)

//go:generate mockgen -source=balance.go -destination=mock/currency_mock.go

// CurrencyConvert is a struct for currency conversion.
type CurrencyConvert struct {
	log *slog.Logger
}

// New creates a new instance of CurrencyConvert.
func New(token string, log *slog.Logger) *CurrencyConvert {
	conv := &CurrencyConvert{log: log}
	go conv.updateRubToUSDRate(token)

	return conv
}

// ConverterResponse represents the response from the currency converter API.
type ConverterResponse struct {
	Success bool           `json:"success"`
	Query   ConverterQuery `json:"query"`
	Result  float64        `json:"result"`
}

// ConverterQuery represents the query parameters for currency conversion.
type ConverterQuery struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// updateRubToUSDRate updates the Rub to USD exchange rate periodically.
func (c *CurrencyConvert) updateRubToUSDRate(token string) {
	const op = "currencyconverter.updateRubToUSDRate"

	log := c.log.With(slog.String("op", op))

	const url = "https://api.apilayer.com/fixer/convert?to=RUB&from=USD&amount=1"

	var response ConverterResponse

	var client *http.Client = &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error("Can't initialize request ", "Error", err)
	}

	req.Header.Set("apikey", token)

	for {
		res, err := client.Do(req)
		if res.Body == nil || err != nil {
			log.Error("Can't do request for currency converting")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error("Can't read response for conversion")
		}

		// Unmarshal the response body into the ConverterResponse struct.
		if err := json.Unmarshal(body, &response); err != nil {
			log.Error("Can't unmarshal response body", "error", err)
		}

		rubToUsdRate = response.Result

		// Print the updated rate.
		log.Info(fmt.Sprintf("Rub to USD Rate updated, now rate is %.2f", rubToUsdRate))
		time.Sleep(time.Minute)

		if err := res.Body.Close(); err != nil {
			log.Warn("Can't close convert response body", "error", err)
		}
	}
}

// Convert converts the given amount from one currency to another.
func (c *CurrencyConvert) Convert(amount float64, currency string) (float64, error) {
	if currency == "USD" {
		return strconv.ParseFloat(fmt.Sprintf("%.2f", amount*rubToUsdRate), 64)
	}

	return strconv.ParseFloat(fmt.Sprintf("%.2f", amount/rubToUsdRate), 64)
}
