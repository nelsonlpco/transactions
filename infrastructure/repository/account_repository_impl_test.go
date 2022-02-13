package repository_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
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
	validAccountEntity := entity.NewAccount(validId, validDocument)
	validAccountModel, _ := new(inframodel.AccountModel).FromEntity(validAccountEntity)
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
		expectedError := commonerrors.NewErrorInternalServer("sql", "invalid query")

		accountDatasource.EXPECT().Create(rootCtx, validAccountModel).Return(expectedError)

		err := accountRepository.Create(rootCtx, validAccountEntity)

		require.Equal(t, expectedError, err)
	})
}

func Test_GetAccountByIdHandler(t *testing.T) {
	validDocument := "10094138052"
	validId := uuid.New()
	binaryId, _ := validId.MarshalBinary()

	validAccountEntity := entity.NewAccount(validId, validDocument)
	validAccountModel, _ := new(inframodel.AccountModel).FromEntity(validAccountEntity)
	rootCtx := context.Background()

	t.Run("should be get a valid account by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)

		accountDatasource.EXPECT().GetById(rootCtx, binaryId).Return(validAccountModel, nil)

		account, err := accountRepository.GetById(rootCtx, validId)

		require.Nil(t, err)
		require.Equal(t, validAccountEntity, account)
	})

	t.Run("should be return an error when accountDatasource fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
		expectedError := commonerrors.NewErrorInternalServer("sql", "invalid query")

		accountDatasource.EXPECT().GetById(rootCtx, binaryId).Return(nil, expectedError)

		entity, err := accountRepository.GetById(rootCtx, validId)

		require.Nil(t, entity)
		require.Equal(t, expectedError, err)
	})

	t.Run("should be return an no content error when account not find", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		accountDatasource := mock_datasource.NewMockAccountDatasource(ctrl)
		accountRepository := repository.NewAccountRepositoryImpl(accountDatasource)
		expectedError := commonerrors.NewErrorNoContent("account not found")

		accountDatasource.EXPECT().GetById(rootCtx, binaryId).Return(nil, expectedError)

		entity, err := accountRepository.GetById(rootCtx, validId)

		require.Nil(t, entity)
		require.Equal(t, expectedError, err)
	})
}
