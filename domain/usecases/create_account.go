package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
)

type CreateAccount struct {
	accountRepository repository.AccountRepository
}

func NewCreateAccount(accountRepository repository.AccountRepository) *CreateAccount {
	return &CreateAccount{
		accountRepository: accountRepository,
	}
}

func (c *CreateAccount) Call(ctx context.Context, account *entity.Account) error {
	accountErrors := account.Validate()

	if accountErrors != nil {
		accountError := domainerrors.ErrorsToError(accountErrors)
		return fmt.Errorf("createAccount: %v", accountError)
	}

	err := c.accountRepository.Create(ctx, account)
	if err != nil {
		return fmt.Errorf("createAccount: %v", err)
	}

	return nil
}
