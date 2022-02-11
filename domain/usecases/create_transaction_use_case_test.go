package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func createTransactionUseCaseRepositoriesMock(t *testing.T) (
	*gomock.Controller,
	*mock_repository.MockTransactionRepository) {
	ctrl := gomock.NewController(t)
	transactionRepository := mock_repository.NewMockTransactionRepository(ctrl)

	return ctrl, transactionRepository
}

func Test_should_be_create_a_create_transaction_use_case(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	useCase := usecases.NewCreateTransactionUseCase(transactionRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_create_a_new_valid_transaction(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	id := uuid.New()
	documentNumber := "10094138052"
	account := entity.NewAccount(id, documentNumber)
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)
	transaction := entity.NewTransaction(id, 20.3434, account, operationType, time.Now())

	ctx := context.Background()
	useCase := usecases.NewCreateTransactionUseCase(transactionRepository)

	transactionRepository.EXPECT().Create(ctx, transaction).Return(nil)

	err := useCase.Call(ctx, transaction)

	require.Nil(t, err)
}

func Test_should_be_create_return_error_when_transaction_and_operation_is_nil(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	id := uuid.New()
	transaction := entity.NewTransaction(id, 0, nil, nil, time.Now())

	expectedError := transaction.Validate()

	ctx := context.Background()
	useCase := usecases.NewCreateTransactionUseCase(transactionRepository)

	err := useCase.Call(ctx, transaction)

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError, err)
}
func Test_should_be_create_return_error_when_transaction_is_invalid(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	id := uuid.New()
	documentNumber := "1008052"
	account := entity.NewAccount(id, documentNumber)
	operationType := entity.NewOperationType(id, "", valueobjects.Credit)
	transaction := entity.NewTransaction(id, 0, account, operationType, time.Now())

	expectedError := transaction.Validate()

	ctx := context.Background()
	useCase := usecases.NewCreateTransactionUseCase(transactionRepository)

	err := useCase.Call(ctx, transaction)

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError, err)
}

func Test_should_be_create_return_error_when_transactionRepository_fail(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	ctx := context.Background()
	id := uuid.New()
	documentNumber := "10094138052"
	account := entity.NewAccount(id, documentNumber)
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)
	transaction := entity.NewTransaction(id, 10, account, operationType, time.Now())
	useCase := usecases.NewCreateTransactionUseCase(transactionRepository)
	expectedError := useCase.MakeError("fail")

	transactionRepository.EXPECT().Create(ctx, transaction).Return(errors.New("fail"))

	err := useCase.Call(ctx, transaction)

	var errorInternalServer *domainerrors.ErrorInternalServer

	require.True(t, errors.As(err, &errorInternalServer))
	require.Equal(t, expectedError, err)
}
