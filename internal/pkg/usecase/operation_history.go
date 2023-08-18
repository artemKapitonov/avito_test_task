package usecase

import (
	"context"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
)

type HistoryUseCase struct {
	OperationHistory
}

type OperationHistory interface {
	Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error)
}

func (h *HistoryUseCase) Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error) {
	return h.OperationHistory.Get(ctx, userID, sort, isDesc)
}
