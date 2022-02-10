package repository

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type OperationTypeRepositoryImpl struct {
	operationTypeDatasource datasource.OperationTypeDatasource
}

func NewOperationTypeRepositoryImpl(operationTypeDatasource datasource.OperationTypeDatasource) *OperationTypeRepositoryImpl {
	return &OperationTypeRepositoryImpl{
		operationTypeDatasource: operationTypeDatasource,
	}
}

func (o *OperationTypeRepositoryImpl) Create(ctx context.Context, operationType *entity.OperationType) error {
	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          int(operationType.GetId()),
		Description: operationType.GetDescription(),
		Operation:   byte(operationType.GetOperation()),
	}

	err := o.operationTypeDatasource.Create(ctx, operationTypeModel)
	if err != nil {
		return fmt.Errorf("operationTypeRepositoryImpl: %v", err)
	}

	return nil
}

func (o *OperationTypeRepositoryImpl) GetById(ctx context.Context, id valueobjects.Id) (*entity.OperationType, error) {
	operationModel, err := o.operationTypeDatasource.GetById(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("operationTypeRepositoryImpl: %v", err)
	}

	operationEntity := entity.NewOperationType(valueobjects.Id(operationModel.Id),
		operationModel.Description,
		valueobjects.Operation(operationModel.Operation))

	return operationEntity, nil
}
