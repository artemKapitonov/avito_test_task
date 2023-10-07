package convert

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	rubToUsdRate float64
)

//go:generate mockgen -source=balance.go -destination=mock/currency_mock.go

// CurrencyConvert is a struct for currency conversion.
type CurrencyConvert struct{}

// New creates a new instance of CurrencyConvert.
func New(token string) *CurrencyConvert {
	go updateRubToUSDRate(token)

	return &CurrencyConvert{}
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
func updateRubToUSDRate(token string) {
	// Rubles in 1 Dollar.
	for {
		// Set the URL for the API call.
		url := "https://api.apilayer.com/fixer/convert?to=RUB&from=USD&amount=1"

		var response ConverterResponse

		client := &http.Client{}

		req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
		if err != nil {
			logrus.Fatalf("Can't initialize request Error: %s", err.Error())
		}

		req.Header.Set("apikey", token)

		res, err := client.Do(req)
		if res.Body == nil || err != nil {
			logrus.Fatal("Can't do request for currency converting")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			logrus.Fatal("Can't read response for convertation")
		}

		// Unmarshal the response body into the ConverterResponse struct.
		if err := json.Unmarshal(body, &response); err != nil {
			logrus.Fatalf("Can't unmarshal response body Error: %s", err.Error())
		}

		rubToUsdRate = response.Result

		// Print the updated rate.
		logrus.Printf("Rub to USD Rate updated, now rate is %.2f", rubToUsdRate)
		time.Sleep(time.Minute)

		if err := res.Body.Close(); err != nil {
			logrus.Warnf("Can't close convert response body Error: %s", err.Error())
		}
	}
}

// Convert converts the given amount from one currency to another.
func (c *CurrencyConvert) Convert(amount float64, fromCurrency string) (float64, error) {

	if fromCurrency == "USD" {
		return strconv.ParseFloat(fmt.Sprintf("%.2f", amount*rubToUsdRate), 64)
	}

	return strconv.ParseFloat(fmt.Sprintf("%.2f", amount/rubToUsdRate), 64)
}
