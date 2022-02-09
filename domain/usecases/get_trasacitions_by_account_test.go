package usecases_test

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func createGetTransactionByAccountResources(t *testing.T) (*gomock.Controller, *mock_repository.MockTransactionRepository) {
	ctrl := gomock.NewController(t)
	transactionRepository := mock_repository.NewMockTransactionRepository(ctrl)

	return ctrl, transactionRepository
}

func createRandomTransaction(account *entity.Account) *entity.Transaction {
	operationTypId := rand.Intn(1000) + 1
	transactionId := rand.Intn(1000) + 1
	operation := rand.Intn(1)

	amount := 10 + rand.Float64()*(100000-10)

	operationType := entity.NewOperationType(valueobjects.Id(operationTypId), "OPERACAO", valueobjects.Operation(operation))

	return entity.NewTransaction(valueobjects.Id(transactionId), valueobjects.Money(amount), account, operationType, time.Now())
}

func Test_should_be_create_get_transaction_by_account_use_case(t *testing.T) {
	ctrl, transactionRepository := createGetTransactionByAccountResources(t)
	defer ctrl.Finish()

	useCase := usecases.NewGetTransactionByAccount(transactionRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_get_valid_transactions(t *testing.T) {
	ctrl, transactionRepository := createGetTransactionByAccountResources(t)
	defer ctrl.Finish()

	account := entity.NewAccount(1, "10094138052")

	expectedTransactions := []*entity.Transaction{
		createRandomTransaction(account),
		createRandomTransaction(account),
		createRandomTransaction(account),
		createRandomTransaction(account),
		createRandomTransaction(account),
	}

	ctx := context.Background()
	useCase := usecases.NewGetTransactionByAccount(transactionRepository)

	transactionRepository.EXPECT().GetTransactionsByAccount(ctx, account).Return(expectedTransactions, nil)

	transactions, err := useCase.Call(ctx, account)

	require.NotNil(t, transactions)
	require.Nil(t, err)

	require.Equal(t, expectedTransactions, transactions)
}

func Test_should_be_return_error_when_repositoryTransaction_fail(t *testing.T) {
	ctrl, transactionRepository := createGetTransactionByAccountResources(t)
	defer ctrl.Finish()

	account := entity.NewAccount(1, "10094138052")
	expectedError := fmt.Errorf("getTransactionByAccount: fail")

	ctx := context.Background()
	useCase := usecases.NewGetTransactionByAccount(transactionRepository)

	transactionRepository.EXPECT().GetTransactionsByAccount(ctx, account).Return(nil, errors.New("fail"))

	transactions, err := useCase.Call(ctx, account)

	require.Nil(t, transactions)

	require.Equal(t, expectedError, err)
}
