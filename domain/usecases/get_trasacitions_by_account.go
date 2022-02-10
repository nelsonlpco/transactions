package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type GetTransactionByAccount struct {
	transactionRepository repository.TransactionRepository
}

func NewGetTransactionByAccount(
	transactionRepository repository.TransactionRepository,
) *GetTransactionByAccount {
	return &GetTransactionByAccount{
		transactionRepository: transactionRepository,
	}
}

func (g *GetTransactionByAccount) Call(ctx context.Context, accountId valueobjects.Id) ([]*entity.Transaction, error) {
	transactions, err := g.transactionRepository.GetTransactionsByAccountId(ctx, accountId)

	if err != nil {
		return nil, fmt.Errorf("getTransactionByAccount: %v", err)
	}

	return transactions, nil
}
