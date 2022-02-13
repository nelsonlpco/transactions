package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	getAccountByIdUseCase *usecases.GetAccountByIdUseCase
	createAccountUseCase  *usecases.CreateAccountUseCase
}

func NewAccountService(
	getAccountByIdUseCase *usecases.GetAccountByIdUseCase,
	createAccountUseCase *usecases.CreateAccountUseCase,
) *AccountService {
	return &AccountService{
		getAccountByIdUseCase: getAccountByIdUseCase,
		createAccountUseCase:  createAccountUseCase,
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, account *entity.Account) error {
	if err := a.createAccountUseCase.Call(ctx, account); err != nil {
		logrus.New().WithField("AccountService", "CreateAccount").Error(err)
		return err
	}

	return nil
}

func (a *AccountService) GetAccountById(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	account, err := a.getAccountByIdUseCase.Call(ctx, id)
	if err != nil {
		logrus.New().WithField("AccountService", "GetAccountById").Error(err)
		return nil, err
	}

	return account, nil
}
