package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetTransactionsByAccountId(ctx context.Context, accountId valueobjects.Id) ([]*entity.Transaction, error)
}
