package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_account_repository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
	accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)

	require.NotNil(t, accountRepository)
}

func Test_CreatAccountHandler(t *testing.T) {
	validDocument := "10094138052"
	validId := uuid.New()
	validAccountModel := &inframodel.AccountModel{
		Id:             validId.String(),
		DocumentNumber: validDocument,
	}
	validAccountEntity := entity.NewAccount(validId, validDocument)
	rootCtx := context.Background()

	t.Run("should be create a new valida acount", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)

		accountDatasource.EXPECT().Create(rootCtx, validAccountModel).Return(nil)

		err := accountRepository.Create(rootCtx, validAccountEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an error when accountDatasource fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
		expectedError := accountRepository.MakeError("fail")

		accountDatasource.EXPECT().Create(rootCtx, validAccountModel).Return(errors.New("fail"))

		err := accountRepository.Create(rootCtx, validAccountEntity)

		require.Equal(t, expectedError, err)
	})
}

func Test_GetAccountByIdHandler(t *testing.T) {
	validDocument := "10094138052"
	validId := uuid.New()

	validAccountModel := &inframodel.AccountModel{
		Id:             validId.String(),
		DocumentNumber: validDocument,
	}
	validAccountEntity := entity.NewAccount(validId, validDocument)
	rootCtx := context.Background()

	t.Run("should be get a valid account by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)

		accountDatasource.EXPECT().GetById(rootCtx, validId.String()).Return(validAccountModel, nil)

		account, err := accountRepository.GetById(rootCtx, validId)

		require.Nil(t, err)
		require.Equal(t, validAccountEntity, account)
	})

	t.Run("should be return an error when accountDatasource fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
		expectedError := accountRepository.MakeError("fail")

		accountDatasource.EXPECT().GetById(rootCtx, validId.String()).Return(nil, errors.New("fail"))

		entity, err := accountRepository.GetById(rootCtx, validId)

		require.Nil(t, entity)
		require.Equal(t, expectedError, err)
	})
}
