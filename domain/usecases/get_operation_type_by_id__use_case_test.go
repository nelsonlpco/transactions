package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func createRepositoryMock(t *testing.T) (*gomock.Controller, *mock_repository.MockOperationTypeRepository) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	operationTypeRepository := mock_repository.NewMockOperationTypeRepository(ctrl)

	return ctrl, operationTypeRepository
}

func Test_should_be_create_a_get_operation_type_by_id_use_case(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	useCase := usecases.NewGetOperationTypeByIdUseCase(operationTypeRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_get_a_valid_operation_type_by_id(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	id := uuid.New()
	operationTypeTest := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)

	ctx := context.Background()
	useCase := usecases.NewGetOperationTypeByIdUseCase(operationTypeRepository)

	operationTypeRepository.EXPECT().GetById(ctx, id).Return(operationTypeTest, nil)

	operationType, err := useCase.Call(ctx, id)

	require.NotNil(t, operationType)
	require.Nil(t, err)
}

func Test_should_be_returns_an_error_when_get_operation_type_invalid(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	ctx := context.Background()
	id := uuid.New()
	operationTypeTest := entity.NewOperationType(id, "", valueobjects.Credit)
	useCase := usecases.NewGetOperationTypeByIdUseCase(operationTypeRepository)
	expectedError := useCase.MakeError(operationTypeTest.Validate().Error())

	operationTypeRepository.EXPECT().GetById(ctx, id).Return(operationTypeTest, nil)

	operationType, err := useCase.Call(ctx, id)

	var errorInternalServer *domainerrors.ErrorInternalServer

	require.Nil(t, operationType)
	require.True(t, errors.As(err, &errorInternalServer))
	require.Equal(t, expectedError.Error(), err.Error())
}

func Test_should_be_returns_an_error_when_operationTypeRepository_fail(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	id := uuid.New()
	ctx := context.Background()
	useCase := usecases.NewGetOperationTypeByIdUseCase(operationTypeRepository)
	expectedError := useCase.MakeError("fail")

	operationTypeRepository.EXPECT().GetById(ctx, id).Return(nil, errors.New("fail"))

	operationType, err := useCase.Call(ctx, id)

	var errorInternalServer *domainerrors.ErrorInternalServer

	require.Nil(t, operationType)
	require.True(t, errors.As(err, &errorInternalServer))
	require.Equal(t, expectedError.Error(), err.Error())
}
