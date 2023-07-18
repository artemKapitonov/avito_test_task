package usecase

import "github.com/artemKapitonov/avito_test_task/internal/pkg/usecase/storage"

type  UseCase struct {
	Account
	Balance
	OperationHistory
	//Converter
}

func New(storage *storage.Storage) *UseCase {
	return &UseCase{
		Account:          &storage.Account,
		Balance:          &storage.Balance,
		OperationHistory: &storage.OperationHistory,
		//Converter: converter.Converter
	}
}
