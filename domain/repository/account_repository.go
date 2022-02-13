package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entity.Account) error
	GetById(ctx context.Context, id uuid.UUID) (*entity.Account, error)
	GetByDocumentNumber(ctx context.Context, documentNumber string) (*entity.Account, error)
}
