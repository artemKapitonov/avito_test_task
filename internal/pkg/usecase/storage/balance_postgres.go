package storage

import (
	"context"
	"errors"
	"fmt"
	"math"
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

func (b *Balance) Update(ctx context.Context, userID uint64, amount float64) (err error) {
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

	row := tx.QueryRow(ctx, operationQuery, operationType, math.Abs(amount), createdDT)
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

func (b *Balance) Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error {
	var sendOperationID, receiveOperationID uint64

	createdDT := time.Now()

	senderQuery := fmt.Sprintf("update %s set balance = balance - $1 where id = $2", usersTable)

	recipientQuery := fmt.Sprintf("update %s set balance = balance + $1 where id = $2", usersTable) //FIXME

	senderOperationQuery := fmt.Sprintf(`insert into %s (operation_type, amount, created_dt) values('send', $1, $2) returning id`, operationsTable)

	recipientOperationQuery := fmt.Sprintf(`insert into %s (operation_type, amount, created_dt) values('receive', $1, $2) returning id`, operationsTable)

	usersOperationQuerySend := fmt.Sprintf("insert into %s (user_id, operation_id) values($1, $2)", usersOperationsTable)

	usersOperationQueryReceive := fmt.Sprintf("insert into %s (user_id, operation_id) values($1, $2)", usersOperationsTable)

	tx, err := b.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, senderQuery, amount, senderID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, recipientQuery, amount, recipientID) //FIXME
	if err != nil {
		return err
	}

	senderRow := tx.QueryRow(ctx, senderOperationQuery, math.Abs(amount), createdDT)
	if err := senderRow.Scan(&sendOperationID); err != nil {
		return err
	}

	_, err = tx.Exec(ctx, usersOperationQuerySend, senderID, sendOperationID)
	if err != nil {
		return err
	}

	recipientRow := tx.QueryRow(ctx, recipientOperationQuery, math.Abs(amount), createdDT)
	if err := recipientRow.Scan(&receiveOperationID); err != nil {
		return err
	}

	_, err = tx.Exec(ctx, usersOperationQueryReceive, recipientID, receiveOperationID)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
