package repository

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
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
	transactionModel, err := new(inframodel.TransactionModel).FromEntity(transaction)
	if err != nil {
		return t.MakeError(err.Error())
	}

	err = t.transactionDatasource.Create(ctx, transactionModel)
	if err != nil {
		return err
	}

	return nil
}

func (TransactionRepositoryImpl) MakeError(errorMessage string) error {
	return commonerrors.NewErrorInternalServer("TransactionRepositoryImpl", errorMessage)
}
