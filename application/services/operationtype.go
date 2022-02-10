package services

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type OperationTypeService struct {
	getOperationTypeByIdUseCase *usecases.GetOperationTypeById
	createOperationTypeUseCase  *usecases.CreateOperationType
}

func NewOperationTypeService(
	getOperationTypeByIdUseCase *usecases.GetOperationTypeById,
	createOperationTypeUseCase *usecases.CreateOperationType,
) *OperationTypeService {
	return &OperationTypeService{
		getOperationTypeByIdUseCase: getOperationTypeByIdUseCase,
		createOperationTypeUseCase:  createOperationTypeUseCase,
	}
}

func (o *OperationTypeService) CreateOperationType(ctx context.Context, operationType *entity.OperationType) error {
	err := o.createOperationTypeUseCase.Call(ctx, operationType)
	if err != nil {
		return err
	}

	return nil
}

func (o *OperationTypeService) GetOperationTypeById(ctx context.Context, id valueobjects.Id) (*entity.OperationType, error) {
	operationType, err := o.getOperationTypeByIdUseCase.Call(ctx, id)
	if err != nil {
		return nil, err
	}

	return operationType, nil
}
