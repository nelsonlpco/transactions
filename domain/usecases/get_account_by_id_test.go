package usecases_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_get_account_by_document_number_use_case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewGetAccountById(accountRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_get_a_valid_account(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	documentNumber := "10094138052"
	accountId := valueobjects.Id(1)
	accountTest := entity.NewAccount(1, documentNumber)

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewGetAccountById(accountRepository)

	accountRepository.EXPECT().GetById(ctx, accountId).Return(accountTest, nil)

	account, err := useCase.Call(ctx, accountId)

	require.Nil(t, err)
	require.NotNil(t, account)
}

func Test_should_be_get_an_error_when_get_an_invalid_account(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	documentNumber := "00023993900"
	accountId := valueobjects.Id(1)
	accountTest := entity.NewAccount(accountId, documentNumber)
	expectError := fmt.Errorf("getAccountById: %v", domainerrors.ErrorsToError(accountTest.Validate()))

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewGetAccountById(accountRepository)

	accountRepository.EXPECT().GetById(ctx, accountId).Return(accountTest, nil)

	account, err := useCase.Call(ctx, accountId)

	require.Nil(t, account)
	require.Equal(t, expectError, err)
}

func Test_should_be_get_an_error_when_repository_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountId := valueobjects.Id(1)
	expectError := fmt.Errorf("getAccountById: fail")

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewGetAccountById(accountRepository)

	accountRepository.EXPECT().GetById(ctx, accountId).Return(nil, errors.New("fail"))

	account, err := useCase.Call(ctx, accountId)

	require.Nil(t, account)
	require.Equal(t, expectError, err)
}
