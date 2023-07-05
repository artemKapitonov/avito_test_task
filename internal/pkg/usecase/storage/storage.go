package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	Account
	Balance
	OperationHistory
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		Account{db: db},
		Balance{db: db},
		OperationHistory{db: db},
	}
}
