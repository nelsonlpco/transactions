package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type OperationTypeRepository interface {
	Create(ctx context.Context, operationType *entity.OperationType) error
	GetById(ctx context.Context, id valueobjects.Id) (*entity.OperationType, error)
}
