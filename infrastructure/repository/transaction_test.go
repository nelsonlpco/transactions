package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_a_transaction_repository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)

	transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)

	require.NotNil(t, transactionRepository)
}

func Test_CreateTransactionHandler(t *testing.T) {
	rootCtx := context.Background()
	validDocument := "10094138052"
	validId := 1
	eventDate := time.Now()
	validAccountModel := &inframodel.AccountModel{
		Id:             validId,
		DocumentNumber: validDocument,
	}
	validAccountEntity := entity.NewAccount(valueobjects.Id(validId), validDocument)
	creditOperationEntity := entity.NewOperationType(valueobjects.Id(validId), "PAGAMENTO", valueobjects.Credit)
	creditOperationModel := &inframodel.OperationTypeModel{
		Id:          validId,
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}
	transactionEntity := entity.NewTransaction(1, 100.32123, validAccountEntity, creditOperationEntity, eventDate)
	transactionModel := &inframodel.TransactionModel{
		Id:            1,
		Amount:        100.32123,
		EventDate:     eventDate,
		Account:       validAccountModel,
		OperationType: creditOperationModel,
	}

	t.Run("should be create a new valid transaction", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)
		transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)

		transactionDatasource.EXPECT().Create(rootCtx, transactionModel).Return(nil)

		err := transactionRepository.Create(rootCtx, transactionEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an error when transactionDatasource fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedError := errors.New("transactionRepositoryImpl: fail")
		transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)
		transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)

		transactionDatasource.EXPECT().Create(rootCtx, transactionModel).Return(errors.New("fail"))

		err := transactionRepository.Create(rootCtx, transactionEntity)

		require.Equal(t, expectedError, err)
	})

	t.Run("should be get transactions by document", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionsModel := []*inframodel.TransactionModel{
			{
				Id:            1,
				Amount:        29.32,
				EventDate:     eventDate,
				Account:       validAccountModel,
				OperationType: creditOperationModel,
			},
			{
				Id:            2,
				Amount:        100.32,
				EventDate:     eventDate,
				Account:       validAccountModel,
				OperationType: creditOperationModel,
			},
		}

		expectedTransactions := []*entity.Transaction{
			entity.NewTransaction(
				1,
				29.32,
				validAccountEntity,
				creditOperationEntity,
				eventDate,
			),
			entity.NewTransaction(
				2,
				100.32,
				validAccountEntity,
				creditOperationEntity,
				eventDate,
			),
		}

		transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)
		transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)

		transactionDatasource.EXPECT().GetTransactionsByAccountId(rootCtx, 1).Return(transactionsModel, nil)

		transactions, err := transactionRepository.GetTransactionsByAccountId(rootCtx, validAccountEntity.GetId())

		require.Nil(t, err)
		require.Equal(t, expectedTransactions, transactions)
	})

	t.Run("should be return an error when transactionDatasource fail to get transactions", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedError := errors.New("transactionRepositoryImpl: fail")
		transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)
		transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)

		transactionDatasource.EXPECT().GetTransactionsByAccountId(rootCtx, 1).Return(nil, errors.New("fail"))

		transactions, err := transactionRepository.GetTransactionsByAccountId(rootCtx, validAccountEntity.GetId())

		require.Nil(t, transactions)
		require.Equal(t, expectedError, err)
	})
}
