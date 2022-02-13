package usecases

import (
	"context"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/sirupsen/logrus"
)

type CreateAccountUseCase struct {
	accountRepository repository.AccountRepository
}

func NewCreateAccountUseCase(accountRepository repository.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountRepository: accountRepository,
	}
}

func (c *CreateAccountUseCase) Call(ctx context.Context, account *entity.Account) error {
	err := account.Validate()
	if err != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "CreateAccountUseCase"}).Error(err.Error())
		return err
	}

	err = c.accountRepository.Create(ctx, account)
	if err != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "CreateAccountUseCase"}).Error(err.Error())
		return err
	}

	logrus.WithFields(logrus.Fields{"UseCase": "CreateAccountUseCase"}).Info("success on create account", account.GetId().ID())
	return nil
}
