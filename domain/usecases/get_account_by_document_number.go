package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
)

type GetAccountByDocumentNumber struct {
	accountRepository repository.AccountRepository
}

func NewGetAccountByDocumentNumber(accountRepository repository.AccountRepository) *GetAccountByDocumentNumber {
	return &GetAccountByDocumentNumber{
		accountRepository: accountRepository,
	}
}

func (g *GetAccountByDocumentNumber) Call(ctx context.Context, documentNumber string) (*entity.Account, error) {
	account, err := g.accountRepository.GetByDocumentNumber(ctx, documentNumber)

	if err != nil {
		return nil, fmt.Errorf("getAccountByDocumentNumber: %v", err)
	}

	accountErrors := account.Validate()
	if accountErrors != nil {
		errorMessage := domainerrors.ErrorsToError(accountErrors)
		return nil, fmt.Errorf("getAccountByDocumentNumber: %v", errorMessage)
	}

	return account, nil
}
