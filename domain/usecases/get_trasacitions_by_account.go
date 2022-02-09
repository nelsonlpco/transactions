package usecases

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
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

func (g *GetTransactionByAccount) Call(ctx context.Context, account *entity.Account) ([]*entity.Transaction, error) {
	transactions, err := g.transactionRepository.GetTransactionsByAccount(ctx, account)

	if err != nil {
		return nil, fmt.Errorf("getTransactionByAccount: %v", err)
	}

	return transactions, nil
}
