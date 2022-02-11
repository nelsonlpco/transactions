package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	mock_repository "github.com/nelsonlpco/transactions/domain/repository/mock"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_get_account_by_document_number_use_case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewGetAccountByIdUseCase(accountRepository)

	require.NotNil(t, useCase)
}

func Test_should_be_get_a_valid_account(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	documentNumber := "10094138052"
	accountId := uuid.New()
	accountTest := entity.NewAccount(accountId, documentNumber)

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	useCase := usecases.NewGetAccountByIdUseCase(accountRepository)

	accountRepository.EXPECT().GetById(ctx, accountId).Return(accountTest, nil)

	account, err := useCase.Call(ctx, accountId)

	require.Nil(t, err)
	require.NotNil(t, account)
	require.Equal(t, accountId, account.GetId())
	require.Equal(t, documentNumber, account.GetDocumentNumber())
}

func Test_should_be_returns_an_error_when_get_an_invalid_account(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	documentNumber := "00023993900"
	accountId := uuid.New()
	accountTest := entity.NewAccount(accountId, documentNumber)
	useCase := usecases.NewGetAccountByIdUseCase(accountRepository)
	expectedError := useCase.MakeError(accountTest.Validate().Error())

	accountRepository.EXPECT().GetById(ctx, accountId).Return(accountTest, nil)

	account, err := useCase.Call(ctx, accountId)

	var errorInternalServer *domainerrors.ErrorInternalServer

	require.Nil(t, account)
	require.True(t, errors.As(err, &errorInternalServer))
	require.Equal(t, expectedError.Error(), err.Error())
}

func Test_should_be_get_an_error_when_accountRepository_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepository := mock_repository.NewMockAccountRepository(ctrl)
	ctx := context.Background()
	accountId := uuid.New()
	useCase := usecases.NewGetAccountByIdUseCase(accountRepository)
	expectedError := useCase.MakeError("fail")

	accountRepository.EXPECT().GetById(ctx, accountId).Return(nil, errors.New("fail"))

	account, err := useCase.Call(ctx, accountId)

	var errorInternalServer *domainerrors.ErrorInternalServer

	require.Nil(t, account)
	require.True(t, errors.As(err, &errorInternalServer))
	require.Equal(t, expectedError.Error(), err.Error())
}
