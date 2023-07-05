package storage

import "github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"

type OperationHistory struct {
	db postgresql.Client
}
