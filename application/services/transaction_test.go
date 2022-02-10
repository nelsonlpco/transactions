package services_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func transactionServiceCreateInfra(t *testing.T) (*gomock.Controller, *services.TransactionService, *mock_datasource.MockTransactionDatasource) {
	ctrl := gomock.NewController(t)
	transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)
	transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)
	createTransactionUseCase := usecases.NewCreateTransaction(transactionRepository)
	getTransactionByIdUseCase := usecases.NewGetTransactionByAccount(transactionRepository)
	transactionService := services.NewTransactionService(createTransactionUseCase, getTransactionByIdUseCase)

	return ctrl, transactionService, transactionDatasource
}

func createRandonTransactionByAccountId(accountId valueobjects.Id) (*entity.Transaction, *inframodel.TransactionModel) {
	randId := rand.Intn(10000)
	account := entity.NewAccount(accountId, "91307555063")
	operationType := entity.NewOperationType(valueobjects.Id(randId), "PAGAMENTO", valueobjects.Credit)
	transactionEntity := entity.NewTransaction(valueobjects.Id(randId), valueobjects.Money(float64(randId)/3.5), account, operationType, time.Now())
	transactionModel := &inframodel.TransactionModel{
		Id:        int(transactionEntity.GetId()),
		Amount:    float64(transactionEntity.GetAmount()),
		EventDate: transactionEntity.GetEventDate(),
		Account: &inframodel.AccountModel{
			Id:             int(transactionEntity.GetAccount().GetId()),
			DocumentNumber: transactionEntity.GetAccount().GetDocumentNumber(),
		},
		OperationType: &inframodel.OperationTypeModel{
			Id:          int(transactionEntity.GetOperationType().GetId()),
			Description: transactionEntity.GetOperationType().GetDescription(),
			Operation:   byte(transactionEntity.GetOperationType().GetOperation()),
		},
	}
	return transactionEntity, transactionModel
}

func Test_should_be_create_a_transaction_service(t *testing.T) {
	ctrl, transactionService, _ := transactionServiceCreateInfra(t)
	defer ctrl.Finish()

	require.NotNil(t, transactionService)
}

func Test_CreateTransactionHandler(t *testing.T) {

	t.Run("should be create a transaction", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		transactionEntity, transactionModel := createRandonTransactionByAccountId(1)
		ctx := context.Background()

		transactionDatasource.EXPECT().Create(ctx, transactionModel).Return(nil)

		err := transactionService.CreateTransaction(ctx, transactionEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an error when data source fail on create the transaction", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		transactionEntity, transactionModel := createRandonTransactionByAccountId(1)
		ctx := context.Background()

		transactionDatasource.EXPECT().Create(ctx, transactionModel).Return(errors.New("fail"))

		err := transactionService.CreateTransaction(ctx, transactionEntity)

		require.NotNil(t, err)
	})

	t.Run("should be return an error when receive an invalid transaction", func(t *testing.T) {
		ctrl, transactionService, _ := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		transactionEntity := entity.NewTransaction(0, 0, nil, nil, time.Now())
		ctx := context.Background()

		err := transactionService.CreateTransaction(ctx, transactionEntity)

		require.NotNil(t, err)
	})
}

func Test_GetTransactionsHandler(t *testing.T) {
	t.Run("should be get transactions from account", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		transactionEntity, transactionModel := createRandonTransactionByAccountId(1)
		transactionEntity1, transactionModel1 := createRandonTransactionByAccountId(1)
		transactionEntity2, transactionModel2 := createRandonTransactionByAccountId(1)

		transactionsModel := []*inframodel.TransactionModel{transactionModel, transactionModel1, transactionModel2}
		expectedTransactions := []*entity.Transaction{transactionEntity, transactionEntity1, transactionEntity2}

		ctx := context.Background()

		transactionDatasource.EXPECT().GetTransactionsByAccountId(ctx, 1).Return(transactionsModel, nil)

		transactions, err := transactionService.GetTransactionsByAccountId(ctx, 1)

		require.Nil(t, err)
		require.Equal(t, expectedTransactions, transactions)
	})

	t.Run("should be return an error when data source fail to get transactions by account", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()

		transactionDatasource.EXPECT().GetTransactionsByAccountId(ctx, 1).Return(nil, errors.New("fail"))

		transactions, err := transactionService.GetTransactionsByAccountId(ctx, 1)

		require.Nil(t, transactions)
		require.NotNil(t, err)
	})
}
