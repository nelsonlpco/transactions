package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/sirupsen/logrus"
)

type GetOperationTypeByIdUseCase struct {
	operationTypeRepository repository.OperationTypeRepository
}

func NewGetOperationTypeByIdUseCase(operationTypeRepository repository.OperationTypeRepository) *GetOperationTypeByIdUseCase {
	return &GetOperationTypeByIdUseCase{
		operationTypeRepository: operationTypeRepository,
	}
}

func (g *GetOperationTypeByIdUseCase) Call(ctx context.Context, id uuid.UUID) (*entity.OperationType, error) {
	operationType, err := g.operationTypeRepository.GetById(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "GetOperationTypeById"}).Error(err.Error())
		return nil, g.MakeError(err.Error())
	}

	operationTypeErrors := operationType.Validate()
	if operationTypeErrors != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "GetOperationTypeById"}).Error(operationTypeErrors.Error())
		return nil, g.MakeError(operationTypeErrors.Error())
	}

	logrus.WithFields(logrus.Fields{"UseCase": "GetOperationTypeById"}).Info("success on get operatonType: ", operationType.GetId().String())
	return operationType, nil
}

func (GetOperationTypeByIdUseCase) MakeError(errorMessage string) error {
	return domainerrors.NewErrorInternalServer("GetOperationTypeByIdUseCase", errorMessage)
}
