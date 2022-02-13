package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func transactionServiceCreateInfra(t *testing.T) (
	*gomock.Controller,
	*services.TransactionService,
	*mock_datasource.MockTransactionDatasource,
	*mock_datasource.MockAccountDatasource,
	*mock_datasource.MockOperationTypeDatasource,
) {
	ctrl := gomock.NewController(t)

	accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
	operationTypeDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
	transactionDatasource := mock_datasource.NewMockTransactionDatasource(ctrl)

	accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
	operationTypeRepository := repository.NewOperationTypeRepositoryImpl(operationTypeDatasource)
	transactionRepository := repository.NewTransactionRepositoryImpl(transactionDatasource)

	getAccountByIdUseCase := usecases.NewGetAccountByIdUseCase(accountRepository)
	getOperationTypeByIdUseCase := usecases.NewGetOperationTypeByIdUseCase(operationTypeRepository)
	createTransactionUseCase := usecases.NewCreateTransactionUseCase(transactionRepository)

	transactionService := services.NewTransactionService(
		getAccountByIdUseCase,
		getOperationTypeByIdUseCase,
		createTransactionUseCase,
	)

	return ctrl, transactionService, transactionDatasource, accountDatasource, operationTypeDatasource
}

func Test_should_be_create_a_transaction_service(t *testing.T) {
	ctrl, transactionService, _, _, _ := transactionServiceCreateInfra(t)
	defer ctrl.Finish()

	require.NotNil(t, transactionService)
}

func Test_CreateTransactionHandler(t *testing.T) {
	id := uuid.New()
	binaryId, _ := id.MarshalBinary()
	documentNumber := "91307555063"
	eventDate := time.Now()

	accountEntity := entity.NewAccount(id, documentNumber)
	operationTypeEntity := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)

	accountModel, _ := new(inframodel.AccountModel).FromEntity(accountEntity)
	operationTypeModel, _ := new(inframodel.OperationTypeModel).FromEntity(operationTypeEntity)
	rootCtx := context.Background()

	t.Run("should be create a transaction", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource, accountDatasource, operationTypeDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		operationTypeDatasource.EXPECT().GetById(rootCtx, binaryId).Return(operationTypeModel, nil)
		accountDatasource.EXPECT().GetById(rootCtx, binaryId).Return(accountModel, nil)

		transactionDatasource.EXPECT().Create(rootCtx, gomock.Any()).Return(nil)

		err := transactionService.CreateTransaction(
			rootCtx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		require.Nil(t, err)
	})

	t.Run("should be return an error when data source fail to get an account", func(t *testing.T) {
		ctrl, transactionService, _, accountDatasource, _ := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		expectedError := commonerrors.NewErrorInternalServer("sql", "invalid query")
		accountDatasource.EXPECT().GetById(ctx, binaryId).Return(nil, expectedError)

		err := transactionService.CreateTransaction(
			ctx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		var errorInternalServer *commonerrors.ErrorInternalServer

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})

	t.Run("should be return an error when data source fail to get an operation type", func(t *testing.T) {
		ctrl, transactionService, _, accountDatasource, operationTypeDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, binaryId).Return(accountModel, nil)
		operationTypeDatasource.EXPECT().GetById(ctx, binaryId).Return(nil, commonerrors.NewErrorInternalServer("sql", "invalid query"))

		err := transactionService.CreateTransaction(
			ctx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		var errorInternalServer *commonerrors.ErrorInternalServer

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})

	t.Run("should be return an error when data source fail to create transaction", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource, accountDatasource, operationTypeDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		expectedError := commonerrors.NewErrorNoContent("no content")

		operationTypeDatasource.EXPECT().GetById(ctx, binaryId).Return(operationTypeModel, nil)
		accountDatasource.EXPECT().GetById(ctx, binaryId).Return(accountModel, nil)
		transactionDatasource.EXPECT().Create(ctx, gomock.Any()).Return(expectedError)

		err := transactionService.CreateTransaction(
			ctx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		var errorInternalServer *commonerrors.ErrorNoContent

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})
}
