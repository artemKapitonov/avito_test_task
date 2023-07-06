package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
)

type Balance struct {
	db postgresql.Client
}

func selectOpertionType(amount float64) (string, error) {
	if amount > 0 {
		return "accrual", nil
	} else if amount < 0 {
		return "redeem", nil
	}
	return "", errors.New("amount in transaction is zero")
}

func (b *Balance) Update(ctx context.Context, userID uint64, amount float64) error {
	createdDT := time.Now()
	var operationID uint64

	operationType, err := selectOpertionType(amount)
	if err != nil {
		return err
	}

	balanceQuery := fmt.Sprintf("update %s set balance = balance + $1 where id = $2;", usersTable)
	operationQuery := fmt.Sprintf("insert into %s (operation_type, amount, created_dt) values($1, $2, $3) returning id", operationsTable)
	usersOperationQuery := fmt.Sprintf("insert into %s (user_id, operation_id) values($1, $2)", usersOperationsTable)

	tx, err := b.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, balanceQuery, amount, userID)
	if err != nil {
		return err
	}

	row := tx.QueryRow(ctx, operationQuery, operationType, amount, createdDT)
	if err := row.Scan(&operationID); err != nil {
		return err
	}

	_, err = tx.Exec(ctx, usersOperationQuery, userID, operationID)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
