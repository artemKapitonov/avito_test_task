package storage

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
)

// Balance of user.
type Balance struct {
	db postgresql.Client
}

// NewBalance initialize new Balance struct.
func NewBalance(db postgresql.Client) *Balance {
	return &Balance{
		db: db,
	}
}

// Select type of operation by amount.
func selectOperationType(amount float64) (string, error) {
	if amount > 0 {
		return "accrual", nil
	} else if amount < 0 {
		return "redeem", nil
	}

	return "", errors.New("amount in transaction is zero")
}

// Update accrual or redeem operation with user's balance.
func (b *Balance) Update(ctx context.Context, userID uint64, amount float64) error {
	createdDT := time.Now()

	var operationID uint64

	operationType, err := selectOperationType(amount)
	if err != nil {
		return err
	}

	balanceQuery := fmt.Sprintf("update %s set balance = balance + $1 where id = $2;", usersTable)

	operationQuery := fmt.Sprintf(
		"insert into %s (operation_type, amount, created_dt) values($1, $2, $3) returning id",
		operationsTable,
	)

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

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Transfer amount from users.
func (b *Balance) Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error {
	var err error

	var createdDT = time.Now()

	tx, err := b.db.Begin(ctx)
	if err != nil {
		return err
	}

	err = sendAmount(ctx, tx, amount, senderID)
	if err != nil {
		return err
	}

	err = receiveAmount(ctx, tx, amount, recipientID)
	if err != nil {
		return err
	}

	sendOperationID, err := getSendOperationID(ctx, tx, amount, createdDT)
	if err != nil {
		return err
	}

	usersOperationQuerySend := fmt.Sprintf("insert into %s (user_id, operation_id) values($1, $2)", usersOperationsTable)

	_, err = tx.Exec(ctx, usersOperationQuerySend, senderID, sendOperationID)
	if err != nil {
		return err
	}

	receiveOperationID, err := getReceiveOperationID(ctx, tx, amount, createdDT)
	if err != nil {
		return err
	}

	usersOperationQueryReceive := fmt.Sprintf("insert into %s (user_id, operation_id) values($1, $2)",
		usersOperationsTable)

	_, err = tx.Exec(ctx, usersOperationQueryReceive, recipientID, receiveOperationID)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func sendAmount(ctx context.Context, tx pgx.Tx, amount float64, senderID uint64) error {
	senderQuery := fmt.Sprintf("update %s set balance = balance - $1 where id = $2", usersTable)

	_, err := tx.Exec(ctx, senderQuery, amount, senderID)
	if err != nil {
		return err
	}

	return nil
}

func receiveAmount(ctx context.Context, tx pgx.Tx, amount float64, recipientID uint64) error {
	recipientQuery := fmt.Sprintf("update %s set balance = balance + $1 where id = $2", usersTable)

	_, err := tx.Exec(ctx, recipientQuery, amount, recipientID)
	if err != nil {
		return err
	}

	return nil
}

func getSendOperationID(ctx context.Context, tx pgx.Tx, amount float64, createdDT time.Time) (uint64, error) {
	var sendOperationID uint64

	senderOperationQuery := fmt.Sprintf(`insert into %s (operation_type, amount, created_dt)
	values('send', $1, $2) returning id`,
		operationsTable)

	senderRow := tx.QueryRow(ctx, senderOperationQuery, math.Abs(amount), createdDT)
	if err := senderRow.Scan(&sendOperationID); err != nil {
		return 0, err
	}

	return sendOperationID, nil
}

func getReceiveOperationID(ctx context.Context, tx pgx.Tx, amount float64, createdDT time.Time) (uint64, error) {
	var receiveOperationID uint64

	recipientOperationQuery := fmt.Sprintf(
		`insert into %s (operation_type, amount, created_dt)
				values('receive', $1, $2) returning id`,
		operationsTable)

	recipientRow := tx.QueryRow(ctx, recipientOperationQuery, math.Abs(amount), createdDT)
	if err := recipientRow.Scan(&receiveOperationID); err != nil {
		return 0, err
	}

	return receiveOperationID, nil
}
