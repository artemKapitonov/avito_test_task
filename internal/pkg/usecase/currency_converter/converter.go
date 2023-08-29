package convert

import (
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

type CurrencyConvert struct{}

func New(token string) *CurrencyConvert {
	go updateRubToUSDRate(token)

	return &CurrencyConvert{}
}

type ConverterResponse struct {
	Success bool           `json:"success"`
	Query   ConverterQuery `json:"query"`
	Result  float64        `json:"result"`
}

type ConverterQuery struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func updateRubToUSDRate(token string) {
	//Rubles in 1 Dollar
	for {
		url := "https://api.apilayer.com/fixer/convert?to=RUB&from=USD&amount=1"

		var response ConverterResponse

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("apikey", token)

		if err != nil {
			fmt.Println(err)
		}

		res, err := client.Do(req)
		if res.Body == nil || err != nil {
			logrus.Fatal("Can't do request for currency converting")
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			logrus.Fatal("Can't read response for cunvertation")
		}
		if err := json.Unmarshal(body, &response); err != nil {
			panic("unmarshal")
		}

		rubToUsdRate = response.Result

		logrus.Printf("Rub to Usd Rate updated, now rate is %f.2", rubToUsdRate)
		time.Sleep(time.Minute)
	}
}

func (c *CurrencyConvert) Convert(amount float64, fromCurrency, toCurrency string) (float64, error) {

	if fromCurrency == "USD" {
		return strconv.ParseFloat(fmt.Sprintf("%f2", amount*rubToUsdRate), 64)
	}

	return strconv.ParseFloat(fmt.Sprintf("%.2f", amount/rubToUsdRate), 64)
}
