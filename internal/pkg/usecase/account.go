package usecase

import (
	"context"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
)

//go:generate mockgen -source=account.go -destination=mocks/account_mock.go

// Account represents the interface for working with user accounts.
type Account interface {
	Create(ctx context.Context) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

// AccountUseCase represents the implementation of the Account interface.
type AccountUseCase struct {
	Account
}

// Create creates a new user account.
func (a *AccountUseCase) Create(ctx context.Context) (entity.User, error) {
	return a.Account.Create(ctx)
}

// GetByID returns a user account by ID.
func (a *AccountUseCase) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	return a.Account.GetByID(ctx, id)
}
