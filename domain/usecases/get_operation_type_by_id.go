package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type GetOperationTypeById struct {
	operationTypeRepository repository.OperationTypeRepository
}

func NewGetOperationTypeById(operationTypeRepository repository.OperationTypeRepository) *GetOperationTypeById {
	return &GetOperationTypeById{
		operationTypeRepository: operationTypeRepository,
	}
}

func (g *GetOperationTypeById) Call(ctx context.Context, id valueobjects.Id) (*entity.OperationType, error) {
	operationType, err := g.operationTypeRepository.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getOperationTypeById: %v", err)
	}

	operationTypeErrors := operationType.Validate()
	if operationTypeErrors != nil {
		errorMessage := domainerrors.ErrorsToError(operationTypeErrors)
		return nil, fmt.Errorf("getOperationTypeById: %v", errorMessage)
	}

	return operationType, nil
}
