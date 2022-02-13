package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_create_account_use_case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewCreateAccountUseCase(accountRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_create_an_valid_account(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()
	account := entity.NewAccount(id, "10094138052")

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewCreateAccountUseCase(accountRepository)

	accountRepository.EXPECT().Create(ctx, account).Return(nil)

	err := useCase.Call(ctx, account)

	require.Nil(t, err)
}

func Test_should_be_return_an_error_when_account_is_invalid_on_create_a_count(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()
	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	account := entity.NewAccount(id, "1000000")
	expectedError := account.Validate()

	useCase := usecases.NewCreateAccountUseCase(accountRepository)

	err := useCase.Call(ctx, account)

	var errorInvalidEntity *commonerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}

func Test_should_be_return_an_error_when_accountRepository_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	ctx := context.Background()
	id := uuid.New()
	expectedError := commonerrors.NewErrorInternalServer("SQL", "Error 1179: invalid query")
	account := entity.NewAccount(id, "10094138052")

	useCase := usecases.NewCreateAccountUseCase(accountRepository)

	accountRepository.EXPECT().Create(ctx, account).Return(expectedError)
	err := useCase.Call(ctx, account)

	var errorInternalServer *commonerrors.ErrorInternalServer

	require.True(t, errors.As(err, &errorInternalServer))
	require.Equal(t, expectedError.Error(), err.Error())
}
