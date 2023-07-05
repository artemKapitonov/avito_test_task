package storage

import (
	"context"
	"time"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
)

type Account struct {
	db postgresql.Client
}

func (a *Account) Create(ctx context.Context) (entity.User, error) {
	var user entity.User

	var id, balance uint64

	createdDT := time.Now()

	row := a.db.QueryRow(ctx, "insert  into users (balance, created_dt) values (0, $1) returning *;", createdDT)

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

func (a *Account) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User

	var balance uint64

	var createdDT time.Time

	row := a.db.QueryRow(ctx, "select * from users where id = $1;", id)

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
