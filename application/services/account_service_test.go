package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

var accountServiceDocument = "91307555063"

func accountServiceCreateInfra(t *testing.T) (*gomock.Controller, *services.AccountService, *mock_datasource.MockAccountDatasource) {
	ctrl := gomock.NewController(t)
	accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
	accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
	createAccountUseCase := usecases.NewCreateAccountUseCase(accountRepository)
	getAccountByIdUseCase := usecases.NewGetAccountByIdUseCase(accountRepository)
	accountService := services.NewAccountService(getAccountByIdUseCase, createAccountUseCase)

	return ctrl, accountService, accountDatasource
}

func Test_Should_be_create_an_account_service(t *testing.T) {
	ctrl, accountService, _ := accountServiceCreateInfra(t)
	defer ctrl.Finish()

	require.NotNil(t, accountService)
}

func Test_AccountServiceCreateHandler(t *testing.T) {
	id := uuid.New()
	t.Run("should be create an account", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             id.String(),
			DocumentNumber: accountServiceDocument,
		}
		accountEntity := accountModel.ToEntity()

		ctx := context.Background()

		accountDatasource.EXPECT().Create(ctx, accountModel).Return(nil)

		err := accountService.CreateAccount(ctx, accountEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an invalid entity error when receive an invalid account", func(t *testing.T) {
		ctrl, accountService, _ := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             id.String(),
			DocumentNumber: "11100000",
		}
		accountEntity := accountModel.ToEntity()

		ctx := context.Background()

		err := accountService.CreateAccount(ctx, accountEntity)

		var errorInvalidEntity *domainerrors.ErrorInvalidEntity

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInvalidEntity))
	})

	t.Run("should be return an internal server error when fail data source", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             id.String(),
			DocumentNumber: accountServiceDocument,
		}
		accountEntity := accountModel.ToEntity()

		ctx := context.Background()

		accountDatasource.EXPECT().Create(ctx, accountModel).Return(errors.New("fail"))

		err := accountService.CreateAccount(ctx, accountEntity)

		var errorInternalServer *domainerrors.ErrorInternalServer

		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})
}

func Test_AccountServiceGetAccountHandler(t *testing.T) {
	t.Run("should be get an account by id", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()
		id := uuid.New()
		accountModel := &inframodel.AccountModel{
			Id:             id.String(),
			DocumentNumber: accountServiceDocument,
		}
		accountEntity := entity.NewAccount(id, accountServiceDocument)
		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(accountModel, nil)

		account, err := accountService.GetAccountById(ctx, id)

		require.Nil(t, err)
		require.Equal(t, accountEntity, account)
	})

	t.Run("should be return an error when data source fail on get account by id", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		id := uuid.New()
		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(nil, errors.New("fail"))

		account, err := accountService.GetAccountById(ctx, id)

		require.Nil(t, account)
		require.NotNil(t, err)
	})

	t.Run("should be return an error when data source returns an invalid account", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		id := uuid.New()
		accountModel := &inframodel.AccountModel{
			Id:             id.String(),
			DocumentNumber: "000010230",
		}
		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, id.String()).Return(accountModel, nil)

		account, err := accountService.GetAccountById(ctx, id)

		var errorInternalServer *domainerrors.ErrorInternalServer

		require.Nil(t, account)
		require.NotNil(t, err)
		require.True(t, errors.As(err, &errorInternalServer))
	})
}
