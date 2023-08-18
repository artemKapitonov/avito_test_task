package usecase

import (
	"context"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
)

type AccountUseCase struct {
	Account
}

type Account interface {
	Create(ctx context.Context) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

func (a *AccountUseCase) Create(ctx context.Context) (entity.User, error) {
	return a.Account.Create(ctx)
}

func (a *AccountUseCase) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	return a.Account.GetByID(ctx, id)
}
