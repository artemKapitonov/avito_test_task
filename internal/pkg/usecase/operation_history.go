package usecase

import (
	"context"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
)

// HistoryUseCase is a use case for operation history.
type HistoryUseCase struct {
	OperationHistory
}

//go:generate mockgen -source=operation_history.go -destination=mocks/operation_hisory_mock.go

// OperationHistory is an interface for getting operation history.
type OperationHistory interface {
	Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error)
}

// Get returns the operation history for a user.
func (h *HistoryUseCase) Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error) {
	return h.OperationHistory.Get(ctx, userID, sort, isDesc)
}
