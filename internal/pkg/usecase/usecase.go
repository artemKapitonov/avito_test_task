package usecase

import (
	converter "github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/currency_converter"
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/storage"
)

type UseCase struct {
	Account
	Balance
	OperationHistory
	CurrencyConverter
}

func New(storage *storage.Storage, converter *converter.CurrencyConvert) *UseCase {
	return &UseCase{
		Account:           &storage.Account,
		Balance:           &storage.Balance,
		OperationHistory:  &storage.OperationHistory,
		CurrencyConverter: converter,
	}
}
