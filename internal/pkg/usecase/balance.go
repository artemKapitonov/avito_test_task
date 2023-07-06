package usecase

import "context"

type BalanceUseCase struct {
	Balance
}

type Balance interface {
	Update(ctx context.Context, userID uint64, amount float64) (error)
}

func (b *BalanceUseCase) Update(ctx context.Context, userID uint64, amount float64) (error) {
	return b.Balance.Update(ctx, userID, amount)
}
