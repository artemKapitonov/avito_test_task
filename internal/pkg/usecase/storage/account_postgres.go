package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
)

type Account struct {
	db postgresql.Client
}

func (a *Account) Create(ctx context.Context) (entity.User, error) {
	var user entity.User

	var id uint64

	createdDT := time.Now()

	query := fmt.Sprintf("insert  into %s (created_dt) values ($1) returning id, created_dt;", usersTable)

	row := a.db.QueryRow(ctx, query, createdDT)

	if err := row.Scan(&id, &createdDT); err != nil {
		return user, err
	}

	user = entity.User{
		ID:        id,
		CreatedDT: createdDT,
	}
	return user, nil
}

func (a *Account) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User

	var balance float64

	var createdDT time.Time

	query := fmt.Sprintf("select * from %s where id = $1;", usersTable)

	row := a.db.QueryRow(ctx, query, id)

	if err := row.Scan(&id, &balance, &createdDT); err != nil {
		return user, err
	}

	user = entity.User{
		ID:        id,
		Balance:   balance,
		CreatedDT: createdDT,
	}

	return user, nil
}
