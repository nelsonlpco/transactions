package services

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type AccountService struct {
	getAccountByIdUseCase *usecases.GetAccountById
	createAccountUseCase  *usecases.CreateAccount
}

func NewAccountService(
	getAccountByIdUseCase *usecases.GetAccountById,
	createAccountUseCase *usecases.CreateAccount,
) *AccountService {
	return &AccountService{
		getAccountByIdUseCase: getAccountByIdUseCase,
		createAccountUseCase:  createAccountUseCase,
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, account *entity.Account) error {
	err := a.createAccountUseCase.Call(ctx, account)

	return err
}

func (a *AccountService) GetAccountById(ctx context.Context, id valueobjects.Id) (*entity.Account, error) {
	account, err := a.getAccountByIdUseCase.Call(ctx, id)
	if err != nil {
		return nil, err
	}

	return account, nil
}
