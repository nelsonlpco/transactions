package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_create_account_use_case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewCreateAccount(accountRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_create_an_valid_account(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := entity.NewAccount(1, "10094138052")

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewCreateAccount(accountRepository)

	accountRepository.EXPECT().Create(ctx, account).Return(nil)

	err := useCase.Call(ctx, account)

	require.Nil(t, err)
}

func Test_should_be_return_an_error_when_account_is_invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := entity.NewAccount(1, "1000000")

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewCreateAccount(accountRepository)

	err := useCase.Call(ctx, account)

	require.NotNil(t, err)
}

func Test_should_be_return_an_error_when_repository_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectError := errors.New("fail")

	account := entity.NewAccount(1, "10094138052")

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewCreateAccount(accountRepository)

	accountRepository.EXPECT().Create(ctx, account).Return(expectError)
	err := useCase.Call(ctx, account)

	require.NotNil(t, err)
}
