package usecase

type CurrencyConvert struct {
	CurrencyConverter
}

type CurrencyConverter interface {
	Convert(amount float64, fromCurrency, toCurrency string) (float64, error)
}

func (c *CurrencyConvert) Convert(amount float64, fromCurrency, toCurrency string) (float64, error) {
	return c.CurrencyConverter.Convert(amount, fromCurrency, toCurrency)
}
