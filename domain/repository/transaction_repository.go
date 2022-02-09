package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetTransactionsByAccount(ctx context.Context, account *entity.Account) ([]*entity.Transaction, error)
}
