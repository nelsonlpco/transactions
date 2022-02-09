package usecases_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
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

	useCase := usecases.NewGetOperationTypeById(operationTypeRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_get_a_valid_operation_type_by_id(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	id := valueobjects.Id(1)
	operationTypeTest := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)

	ctx := context.Background()
	useCase := usecases.NewGetOperationTypeById(operationTypeRepository)

	operationTypeRepository.EXPECT().GetById(ctx, id).Return(operationTypeTest, nil)

	operationType, err := useCase.Call(ctx, id)

	require.NotNil(t, operationType)
	require.Nil(t, err)
}

func Test_should_be_get_an_error_when_operation_type_is_invalid(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	id := valueobjects.Id(1)
	operationTypeTest := entity.NewOperationType(id, "", valueobjects.Credit)
	expectedError := fmt.Errorf("getOperationTypeById: %v", domainerrors.ErrorsToError(operationTypeTest.Validate()))

	ctx := context.Background()
	useCase := usecases.NewGetOperationTypeById(operationTypeRepository)

	operationTypeRepository.EXPECT().GetById(ctx, id).Return(operationTypeTest, nil)

	operationType, err := useCase.Call(ctx, id)

	require.Nil(t, operationType)
	require.Equal(t, expectedError, err)
}

func Test_should_be_get_an_error_when_operationTypeRepositoryFail(t *testing.T) {
	ctrl, operationTypeRepository := createRepositoryMock(t)
	defer ctrl.Finish()

	id := valueobjects.Id(1)
	expectedError := fmt.Errorf("getOperationTypeById: fail")

	ctx := context.Background()
	useCase := usecases.NewGetOperationTypeById(operationTypeRepository)

	operationTypeRepository.EXPECT().GetById(ctx, id).Return(nil, errors.New("fail"))

	operationType, err := useCase.Call(ctx, id)

	require.Nil(t, operationType)
	require.Equal(t, expectedError, err)
}
