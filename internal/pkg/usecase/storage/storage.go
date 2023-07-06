package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	Account
	Balance
	OperationHistory
}

const (
	usersTable           = "users"
	operationsTable      = "operations"
	usersOperationsTable = "user_operations"
)

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		Account{db: db},
		Balance{db: db},
		OperationHistory{db: db},
	}
}
