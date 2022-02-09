package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
)

type CreateTransaction struct {
	transactionRepository repository.TransactionRepository
}

func NewCreateTransaction(
	transactionRepository repository.TransactionRepository,
) *CreateTransaction {
	return &CreateTransaction{
		transactionRepository: transactionRepository,
	}
}

func (c *CreateTransaction) Call(ctx context.Context, transaction *entity.Transaction) error {
	transactionErrors := transaction.Validate()

	if transactionErrors != nil {
		return fmt.Errorf("createTransaction: %v", domainerrors.ErrorsToError(transactionErrors))
	}

	err := c.transactionRepository.Create(ctx, transaction)
	if err != nil {
		return fmt.Errorf("createTransaction: %v", err)
	}

	return nil
}
