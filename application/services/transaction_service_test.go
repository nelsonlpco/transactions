package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
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
	documentNumber := "91307555063"
	eventDate := time.Now()
	accountModel := &inframodel.AccountModel{
		Id:             id.String(),
		DocumentNumber: documentNumber,
	}
	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          id.String(),
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}
	transactionModel := &inframodel.TransactionModel{
		Id:            id.String(),
		Amount:        123.23,
		EventDate:     eventDate,
		Account:       accountModel,
		OperationType: operationTypeModel,
	}

	t.Run("should be create a transaction", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource, accountDatasource, operationTypeDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()

		operationTypeDatasource.EXPECT().GetById(ctx, id.String()).Return(operationTypeModel, nil)
		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(accountModel, nil)
		transactionDatasource.EXPECT().Create(ctx, transactionModel).Return(nil)

		err := transactionService.CreateTransaction(
			ctx,
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

		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(nil, errors.New("fail"))

		err := transactionService.CreateTransaction(
			ctx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		var errorInternalServer *domainerrors.ErrorInternalServer

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})

	t.Run("should be return an error when data source fail to get an operation type", func(t *testing.T) {
		ctrl, transactionService, _, accountDatasource, operationTypeDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(accountModel, nil)
		operationTypeDatasource.EXPECT().GetById(ctx, id.String()).Return(nil, errors.New("fail"))

		err := transactionService.CreateTransaction(
			ctx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		var errorInternalServer *domainerrors.ErrorInternalServer

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})

	t.Run("should be return an error when data source fail to create transaction", func(t *testing.T) {
		ctrl, transactionService, transactionDatasource, accountDatasource, operationTypeDatasource := transactionServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()

		operationTypeDatasource.EXPECT().GetById(ctx, id.String()).Return(operationTypeModel, nil)
		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(accountModel, nil)
		transactionDatasource.EXPECT().Create(ctx, transactionModel).Return(errors.New("fail"))

		err := transactionService.CreateTransaction(
			ctx,
			id,
			id,
			id,
			123.23,
			eventDate,
		)

		var errorInternalServer *domainerrors.ErrorInternalServer

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})
}
