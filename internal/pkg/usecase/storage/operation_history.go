package storage

import (
	"context"
	"fmt"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
)

type OperationHistory struct {
	db postgresql.Client
}

func (h *OperationHistory) Get(ctx context.Context, userID uint64) ([]entity.Operation, error) {
	var operation entity.Operation
	var operations []entity.Operation

	historyQuery := fmt.Sprintf(`select o.id, o.operation_type, o.amount, o.created_dt from %s o
	inner join %s uo on o.id = uo.operation_id where uo.user_id = $1;`, operationsTable, usersOperationsTable)

	rows, err := h.db.Query(ctx, historyQuery, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&operation.ID, &operation.OperationType, &operation.Amount, &operation.CreatedDT); err != nil {
			return nil, err
		}
		operation.Currency = "₽"

		operations = append(operations, operation)
	}

	return operations, nil
}

 
