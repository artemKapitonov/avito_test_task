package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
	"github.com/jackc/pgx/v5"
)

// Account of user
type Account struct {
	db postgresql.Client
}

// NewAccount initialize new account struct
func NewAccount(db postgresql.Client) *Account {
	return &Account{
		db: db,
	}
}

// Create new user in database
func (a *Account) Create(ctx context.Context) (entity.User, error) {
	var user entity.User

	var id uint64

	createdDT := time.Now()

	// Construct the query
	query := fmt.Sprintf("INSERT INTO %s (created_dt) VALUES ($1) RETURNING id, created_dt;", usersTable)

	row := a.db.QueryRow(ctx, query, createdDT)

	if err := row.Scan(&id, &createdDT); err != nil {
		return user, err
	}

	// Create the user object
	user = entity.User{
		ID:        id,
		CreatedDT: createdDT,
	}
	return user, nil
}

// GetByID select user from database by userID
func (a *Account) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User
	var balance float64
	var createdDT time.Time

	// Construct the query
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", usersTable)

	// Execute the query and retrieve the result
	row := a.db.QueryRow(ctx, query, id)

	if err := row.Scan(&id, &balance, &createdDT); err != nil {
		if err == pgx.ErrNoRows {
			return user, errors.New("user with this ID not found")
		}

		return user, err
	}

	// Create the user object
	user = entity.User{
		ID:        id,
		Balance:   balance,
		CreatedDT: createdDT,
	}

	return user, nil
}
