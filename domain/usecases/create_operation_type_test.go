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

func Test_should_be_create_a_create_document_type_use_case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	operationTypeRepository := mock_repository.NewMockOperationTypeRepository(ctrl)
	usecase := usecases.NewCreateOperationType(operationTypeRepository)

	require.NotNil(t, usecase)
}

func Test_should_be_create_a_new_document_type(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	operationTypeTest := entity.NewOperationType(1, "PAGAMENTO", valueobjects.NewOperation(valueobjects.Credit))

	operationTypeRepository := mock_repository.NewMockOperationTypeRepository(ctrl)
	usecase := usecases.NewCreateOperationType(operationTypeRepository)

	operationTypeRepository.EXPECT().Create(ctx, operationTypeTest).Return(nil)

	err := usecase.Call(ctx, operationTypeTest)

	require.Nil(t, err)
}

func Test_should_be_return_an_error_when_operationType_is_invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	operationTypeTest := entity.NewOperationType(1, "PAGAMENTO", valueobjects.NewOperation(2))
	expectedError := fmt.Errorf("createOperationType: %v", domainerrors.ErrorsToError(operationTypeTest.Validate()))

	operationTypeRepository := mock_repository.NewMockOperationTypeRepository(ctrl)
	usecase := usecases.NewCreateOperationType(operationTypeRepository)

	err := usecase.Call(ctx, operationTypeTest)

	require.Equal(t, err, expectedError)
}

func Test_should_be_return_an_error_when_operationTypeRepository_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	operationTypeTest := entity.NewOperationType(1, "PAGAMENTO", valueobjects.NewOperation(valueobjects.Credit))
	expectedError := fmt.Errorf("createOperationType: fail")

	operationTypeRepository := mock_repository.NewMockOperationTypeRepository(ctrl)
	usecase := usecases.NewCreateOperationType(operationTypeRepository)

	operationTypeRepository.EXPECT().Create(ctx, operationTypeTest).Return(errors.New("fail"))

	err := usecase.Call(ctx, operationTypeTest)

	require.Equal(t, expectedError, err)
}
