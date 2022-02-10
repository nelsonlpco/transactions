package repository

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
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
	accountModel := &inframodel.AccountModel{
		Id:             int(transaction.GetAccount().GetId()),
		DocumentNumber: transaction.GetAccount().GetDocumentNumber(),
	}

	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          int(transaction.GetOperationType().GetId()),
		Description: transaction.GetOperationType().GetDescription(),
		Operation:   byte(transaction.GetOperationType().GetOperation()),
	}

	transactionModel := &inframodel.TransactionModel{
		Id:            int(transaction.GetId()),
		Amount:        float64(transaction.GetAmount()),
		EventDate:     transaction.GetEventDate(),
		Account:       accountModel,
		OperationType: operationTypeModel,
	}
	err := t.transactionDatasource.Create(ctx, transactionModel)

	if err != nil {
		return fmt.Errorf("transactionRepositoryImpl: %v", err)
	}

	return nil
}

func (t *TransactionRepositoryImpl) GetTransactionsByAccountId(ctx context.Context, accountId valueobjects.Id) ([]*entity.Transaction, error) {
	transactionsModel, err := t.transactionDatasource.GetTransactionsByAccountId(ctx, int(accountId))
	if err != nil {
		return nil, fmt.Errorf("transactionRepositoryImpl: %v", err)
	}

	var transactions []*entity.Transaction

	for _, transactionModel := range transactionsModel {
		account := entity.NewAccount(valueobjects.Id(transactionModel.Account.Id), transactionModel.Account.DocumentNumber)
		operationType := entity.NewOperationType(
			valueobjects.Id(transactionModel.OperationType.Id),
			transactionModel.OperationType.Description,
			valueobjects.Operation(transactionModel.OperationType.Operation))
		transaction := entity.NewTransaction(
			valueobjects.Id(transactionModel.Id),
			valueobjects.Money(transactionModel.Amount),
			account,
			operationType,
			transactionModel.EventDate)

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
