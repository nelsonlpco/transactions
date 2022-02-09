package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entity.Account) error
	GetByDocumentNumber(ctx context.Context, documentNumber string) (*entity.Account, error)
}
