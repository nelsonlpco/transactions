package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
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
		Id:          operationType.GetId().String(),
		Description: operationType.GetDescription(),
		Operation:   byte(operationType.GetOperation()),
	}

	err := o.operationTypeDatasource.Create(ctx, operationTypeModel)
	if err != nil {
		return o.MakeError(err.Error())
	}

	return nil
}

func (o *OperationTypeRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.OperationType, error) {
	byteId, err := id.MarshalBinary()
	if err != nil {
		return nil, o.MakeError(err.Error())
	}

	operationModel, err := o.operationTypeDatasource.GetById(ctx, byteId)
	if err != nil {
		return nil, o.MakeError(err.Error())
	}

	operationEntity := operationModel.ToEntity()

	return operationEntity, nil
}

func (OperationTypeRepositoryImpl) MakeError(errorMessage string) error {
	return domainerrors.NewErrorInternalServer("OperationRepositoryImpl", errorMessage)
}
