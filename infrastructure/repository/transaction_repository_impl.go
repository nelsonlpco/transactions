package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type TransactionRepositoryImpl struct {
	transactionDatasource datasource.TransactionDatasource
}

func NewTransactionRepositoryImpl(transactionDatasource datasource.TransactionDatasource) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		transactionDatasource: transactionDatasource,
	}
}

func (t *TransactionRepositoryImpl) Create(ctx context.Context, transaction *entity.Transaction) error {
	transactionModel := new(inframodel.TransactionModel).FromEntity(transaction)

	err := t.transactionDatasource.Create(ctx, transactionModel)
	if err != nil {
		return t.MakeError(err.Error())
	}

	return nil
}

func (TransactionRepositoryImpl) MakeError(errorMessage string) error {
	return domainerrors.NewErrorInternalServer("TransactionRepositoryImpl", errorMessage)
}
