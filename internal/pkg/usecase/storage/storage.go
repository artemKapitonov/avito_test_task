package storage

import "github.com/jackc/pgx/v5/pgxpool"

// Storage is postgres database.
type Storage struct {
	*Account
	*Balance
	*OperationHistory
}

const (
	usersTable           = "users"
	operationsTable      = "operations"
	usersOperationsTable = "user_operations"
)

// New storage.
func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		Account:          NewAccount(db),
		Balance:          NewBalance(db),
		OperationHistory: NewOperationHistory(db),
	}
}

// Close is close all db connections.
func (s *Storage) Close() {
	s.Account.db.Close()
	s.Balance.db.Close()
	s.OperationHistory.db.Close()
}
