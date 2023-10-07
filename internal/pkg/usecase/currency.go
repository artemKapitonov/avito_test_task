package usecase

// CurrencyConvert is a struct that implements CurrencyConverter interface.
type CurrencyConvert struct {
	CurrencyConverter
}

//go:generate mockgen -source=currency.go -destination=mocks/currency_mock.go

// CurrencyConverter is an interface that defines the Convert method.
type CurrencyConverter interface {
	Convert(amount float64, fromCurrency string) (float64, error)
}

// Convert is a method of CurrencyConvert that calls the Convert method of CurrencyConverter interface.
func (c *CurrencyConvert) Convert(amount float64, fromCurrency, toCurrency string) (float64, error) {
	return c.CurrencyConverter.Convert(amount, fromCurrency)
}
