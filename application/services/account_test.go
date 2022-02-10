package services_test

import (
	"context"
	"errors"
	"testing"

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

var accountServiceDocument = "91307555063"

func accountServiceCreateInfra(t *testing.T) (*gomock.Controller, *services.AccountService, *mock_datasource.MockAccountDatasource) {
	ctrl := gomock.NewController(t)
	accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
	accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
	createAccountUseCase := usecases.NewCreateAccount(accountRepository)
	getAccountByIdUseCase := usecases.NewGetAccountById(accountRepository)
	accountService := services.NewAccountService(getAccountByIdUseCase, createAccountUseCase)

	return ctrl, accountService, accountDatasource
}

func Test_Should_be_create_an_account_service(t *testing.T) {
	ctrl, accountService, _ := accountServiceCreateInfra(t)
	defer ctrl.Finish()

	require.NotNil(t, accountService)
}

func Test_AccountServiceCreateHandler(t *testing.T) {
	t.Run("should be create an account", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             1,
			DocumentNumber: accountServiceDocument,
		}
		accountEntity := entity.NewAccount(1, accountServiceDocument)
		ctx := context.Background()

		accountDatasource.EXPECT().Create(ctx, accountModel).Return(nil)

		err := accountService.CreateAccount(ctx, accountEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an error when fail data source", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             1,
			DocumentNumber: accountServiceDocument,
		}
		accountEntity := entity.NewAccount(1, accountServiceDocument)
		ctx := context.Background()

		accountDatasource.EXPECT().Create(ctx, accountModel).Return(errors.New("fail"))

		err := accountService.CreateAccount(ctx, accountEntity)

		require.NotNil(t, err)
	})

	t.Run("should be return an error when receive an invalid account", func(t *testing.T) {
		ctrl, accountService, _ := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountEntity := entity.NewAccount(1, "000011111")
		ctx := context.Background()

		err := accountService.CreateAccount(ctx, accountEntity)

		require.NotNil(t, err)
	})
}

func Test_AccountServiceGetAccountHandler(t *testing.T) {
	t.Run("should be get an account by id", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             1,
			DocumentNumber: accountServiceDocument,
		}
		accountEntity := entity.NewAccount(1, accountServiceDocument)
		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, 1).Return(accountModel, nil)

		account, err := accountService.GetAccountById(ctx, 1)

		require.Nil(t, err)
		require.Equal(t, accountEntity, account)
	})

	t.Run("should be return an error when data source fail on get account by id", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, 1).Return(nil, errors.New("fail"))

		account, err := accountService.GetAccountById(ctx, valueobjects.NewId(1))

		require.Nil(t, account)
		require.NotNil(t, err)
	})

	t.Run("should be return an error when data source returns an invalid account", func(t *testing.T) {
		ctrl, accountService, accountDatasource := accountServiceCreateInfra(t)
		defer ctrl.Finish()

		accountModel := &inframodel.AccountModel{
			Id:             1,
			DocumentNumber: "000010230",
		}
		ctx := context.Background()

		accountDatasource.EXPECT().GetById(ctx, 1).Return(accountModel, nil)

		account, err := accountService.GetAccountById(ctx, 1)

		require.Nil(t, account)
		require.NotNil(t, err)
	})
}
