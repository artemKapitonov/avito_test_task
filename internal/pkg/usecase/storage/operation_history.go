package storage

import (
	"context"
	"fmt"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"

	"github.com/artemKapitonov/avito_test_task/pkg/client/postgresql"
)

// OperationHistory of user.
type OperationHistory struct {
	db postgresql.Client
}

// NewOperationHistory initialize new OperationHistory struct.
func NewOperationHistory(db postgresql.Client) *OperationHistory {
	return &OperationHistory{
		db: db,
	}
}

// Get a history of operation by userID with sort and desc params.
func (h *OperationHistory) Get(ctx context.Context, userID uint64, sort string, isDesc bool) ([]entity.Operation, error) {

	var operation entity.Operation

	var operations []entity.Operation

	var historyQuery string

	if sort == "date" {
		historyQuery = fmt.Sprintf(`select o.id, o.operation_type, o.amount, o.created_dt from %s o inner join %s uo on o.id = uo.operation_id
			where uo.user_id = $1 order by o.created_dt desc;`, operationsTable, usersOperationsTable)
	} else {
		if isDesc {
			historyQuery = fmt.Sprintf(
				`select o.id, o.operation_type, o.amount, o.created_dt from %s o inner join %s uo on o.id = uo.operation_id
				where uo.user_id = $1 order by o.amount desc`,
				operationsTable, usersOperationsTable,
			)
		} else {
			historyQuery = fmt.Sprintf(
				`select o.id, o.operation_type, o.amount, o.created_dt from %s o inner join %s uo on o.id = uo.operation_id
			where uo.user_id = $1 order by o.amount asc`,
				operationsTable, usersOperationsTable,
			)
		}
	}

	rows, err := h.db.Query(ctx, historyQuery, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&operation.ID, &operation.OperationType, &operation.Amount, &operation.CreatedDT); err != nil {
			return nil, err
		}

		operation.Currency = "RUB"

		operations = append(operations, operation)
	}

	return operations, nil
}
