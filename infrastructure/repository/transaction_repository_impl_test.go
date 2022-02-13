package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
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
	validId := uuid.New()
	eventDate := time.Now()

	accountEntity := entity.NewAccount(validId, validDocument)
	creditOperationEntity := entity.NewOperationType(validId, "PAGAMENTO", valueobjects.Credit)
	transactionEntity := entity.NewTransaction(validId, 100.32123, accountEntity, creditOperationEntity, eventDate)

	transactionModel, _ := new(inframodel.TransactionModel).FromEntity(transactionEntity)

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

		transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)
		transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)
		expectedError := commonerrors.NewErrorInternalServer("sql", "invalid query")

		transactionDatasource.EXPECT().Create(rootCtx, transactionModel).Return(expectedError)

		err := transactionRepository.Create(rootCtx, transactionEntity)

		require.Equal(t, expectedError, err)
	})
}
