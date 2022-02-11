package usecases

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/sirupsen/logrus"
)

type CreateTransactionUseCase struct {
	transactionRepository repository.TransactionRepository
}

func NewCreateTransactionUseCase(
	transactionRepository repository.TransactionRepository,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionRepository: transactionRepository,
	}
}

func (c *CreateTransactionUseCase) Call(ctx context.Context, transaction *entity.Transaction) error {
	transactionErrors := transaction.Validate()
	if transactionErrors != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "CreateTransactionUseCase"}).Error(transactionErrors.Error())
		return transactionErrors
	}

	err := c.transactionRepository.Create(ctx, transaction)
	if err != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "CreateTransactionUseCase"}).Error(err.Error())
		return c.MakeError(err.Error())
	}

	logrus.WithFields(logrus.Fields{"UseCase": "CreateTransactionUseCase"}).Info("success on create transaction", transaction.GetId().ID())
	return nil
}

func (CreateTransactionUseCase) MakeError(errorMessage string) error {
	return domainerrors.NewErrorInternalServer("CreateTransactionUseCase", errorMessage)
}
