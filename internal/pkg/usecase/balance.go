package usecase

import "context"

type BalanceUseCase struct {
	Balance
}

type Balance interface {
	Update(ctx context.Context, userID uint64, amount float64) error
	Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error
}

func (b *BalanceUseCase) Update(ctx context.Context, userID uint64, amount float64) error {
	return b.Balance.Update(ctx, userID, amount)
}

func (b *BalanceUseCase) Transfer(ctx context.Context, senderID, recipientID uint64, amount float64) error {
	return b.Balance.Transfer(ctx, senderID, recipientID, amount)
}
