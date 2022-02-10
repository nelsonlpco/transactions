package services

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type TransactionService struct {
	createTransactionUseCase       *usecases.CreateTransaction
	getTransactionByAccountUseCase *usecases.GetTransactionByAccount
}

func NewTransactionService(
	createTransactionUseCase *usecases.CreateTransaction,
	getTransactionByAccountUseCase *usecases.GetTransactionByAccount,
) *TransactionService {
	return &TransactionService{
		createTransactionUseCase:       createTransactionUseCase,
		getTransactionByAccountUseCase: getTransactionByAccountUseCase,
	}
}

func (t *TransactionService) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	err := t.createTransactionUseCase.Call(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionService) GetTransactionsByAccountId(ctx context.Context, accountId valueobjects.Id) ([]*entity.Transaction, error) {
	transactions, err := t.getTransactionByAccountUseCase.Call(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
