package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
)

type OperationTypeRepository interface {
	Create(ctx context.Context, operationType *entity.OperationType) error
	GetById(ctx context.Context, id uuid.UUID) (*entity.OperationType, error)
}
