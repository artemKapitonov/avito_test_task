package usecase

import (
	"context"
)

// BalanceUseCase represents the use case for managing balance.
type BalanceUseCase struct {
	Balance
}

//go:generate mockgen -source=balance.go -destination=mocks/balance_mock.go

// Balance is an interface for managing balance.
type Balance interface {
	Update(ctx context.Context, userID uint64, amount float64) error
	Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error
}

// Update updates the balance for a specific user.
func (b *BalanceUseCase) Update(ctx context.Context, userID uint64, amount float64) error {
	return b.Balance.Update(ctx, userID, amount)
}

// Transfer transfers balance from one user to another.
func (b *BalanceUseCase) Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error {
	return b.Balance.Transfer(ctx, senderID, recipientID, amount)
}
