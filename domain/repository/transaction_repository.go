package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
}
