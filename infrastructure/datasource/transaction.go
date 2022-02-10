package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type TransactionDatasource interface {
	Create(ctx context.Context, transaction *inframodel.TransactionModel) error
	GetTransactionsByAccountId(ctx context.Context, accountId int) ([]*inframodel.TransactionModel, error)
}
