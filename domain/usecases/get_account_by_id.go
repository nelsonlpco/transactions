package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type GetAccountById struct {
	accountRepository repository.AccountRepository
}

func NewGetAccountById(accountRepository repository.AccountRepository) *GetAccountById {
	return &GetAccountById{
		accountRepository: accountRepository,
	}
}

func (g *GetAccountById) Call(ctx context.Context, accountId valueobjects.Id) (*entity.Account, error) {
	account, err := g.accountRepository.GetById(ctx, accountId)

	if err != nil {
		return nil, fmt.Errorf("getAccountById: %v", err)
	}

	accountErrors := account.Validate()
	if accountErrors != nil {
		errorMessage := domainerrors.ErrorsToError(accountErrors)
		return nil, fmt.Errorf("getAccountById: %v", errorMessage)
	}

	return account, nil
}
