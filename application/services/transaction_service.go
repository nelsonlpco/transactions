package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type TransactionService struct {
	createTransactionUseCase    *usecases.CreateTransactionUseCase
	getAccountByIdUseCase       *usecases.GetAccountByIdUseCase
	getOperationTypeByIdUseCase *usecases.GetOperationTypeByIdUseCase
}

func NewTransactionService(
	getAccountByIdUseCase *usecases.GetAccountByIdUseCase,
	getOperationTypeByIdUseCase *usecases.GetOperationTypeByIdUseCase,
	createTransactionUseCase *usecases.CreateTransactionUseCase,
) *TransactionService {
	return &TransactionService{
		createTransactionUseCase:    createTransactionUseCase,
		getAccountByIdUseCase:       getAccountByIdUseCase,
		getOperationTypeByIdUseCase: getOperationTypeByIdUseCase,
	}
}

func (t *TransactionService) CreateTransaction(
	ctx context.Context,
	transactionId,
	accountId,
	operationTypeId uuid.UUID,
	amount valueobjects.Money,
	eventDate time.Time,
) error {
	account, err := t.getAccountByIdUseCase.Call(ctx, accountId)
	if err != nil {
		return err
	}

	operationType, err := t.getOperationTypeByIdUseCase.Call(ctx, operationTypeId)
	if err != nil {
		return err
	}

	transaction := entity.NewTransaction(transactionId, amount, account, operationType, eventDate)

	err = t.createTransactionUseCase.Call(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
