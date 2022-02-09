package usecases_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
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

	useCase := usecases.NewCreateTransaction(transactionRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_create_a_new_valid_transaction(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	documentNumber := "10094138052"
	account := entity.NewAccount(1, documentNumber)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)
	transaction := entity.NewTransaction(1, 20.3434, account, operationType, time.Now())

	ctx := context.Background()
	useCase := usecases.NewCreateTransaction(transactionRepository)

	transactionRepository.EXPECT().Create(ctx, transaction).Return(nil)

	err := useCase.Call(ctx, transaction)

	require.Nil(t, err)
}

func Test_should_be_create_return_error_when_transaction_is_invalid(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	documentNumber := "1008052"
	account := entity.NewAccount(1, documentNumber)
	operationType := entity.NewOperationType(0, "PAGAMENTO", valueobjects.Credit)
	transaction := entity.NewTransaction(1, 0, account, operationType, time.Now())

	expectedError := fmt.Errorf("createTransaction: %v", domainerrors.ErrorsToError(transaction.Validate()))

	ctx := context.Background()
	useCase := usecases.NewCreateTransaction(transactionRepository)

	err := useCase.Call(ctx, transaction)

	require.Equal(t, expectedError, err)
}

func Test_should_be_create_return_error_when_transactionRepository_fail(t *testing.T) {
	ctrl, transactionRepository := createTransactionUseCaseRepositoriesMock(t)
	defer ctrl.Finish()

	documentNumber := "10094138052"
	account := entity.NewAccount(1, documentNumber)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)
	transaction := entity.NewTransaction(1, 10, account, operationType, time.Now())
	expectedError := errors.New("createTransaction: fail")

	ctx := context.Background()
	useCase := usecases.NewCreateTransaction(transactionRepository)
	transactionRepository.EXPECT().Create(ctx, transaction).Return(errors.New("fail"))

	err := useCase.Call(ctx, transaction)

	require.Equal(t, expectedError, err)
}
