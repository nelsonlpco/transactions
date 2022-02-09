package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
)

type CreateOperationType struct {
	operationTypeRepository repository.OperationTypeRepository
}

func NewCreateOperationType(operationTypeRepository repository.OperationTypeRepository) *CreateOperationType {
	return &CreateOperationType{
		operationTypeRepository: operationTypeRepository,
	}
}

func (c *CreateOperationType) Call(ctx context.Context, operationType *entity.OperationType) error {
	operationTypeErrors := operationType.Validate()

	if operationTypeErrors != nil {
		return fmt.Errorf("createOperationType: %v", domainerrors.ErrorsToError(operationTypeErrors))
	}

	err := c.operationTypeRepository.Create(ctx, operationType)

	if err != nil {
		return fmt.Errorf("createOperationType: %v", err)
	}

	return nil
}
