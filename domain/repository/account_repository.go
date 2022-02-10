package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entity.Account) error
	GetById(ctx context.Context, id valueobjects.Id) (*entity.Account, error)
}
