package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/repository"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/sirupsen/logrus"
)

type GetAccountByIdUseCase struct {
	accountRepository repository.AccountRepository
}

func NewGetAccountByIdUseCase(accountRepository repository.AccountRepository) *GetAccountByIdUseCase {
	return &GetAccountByIdUseCase{
		accountRepository: accountRepository,
	}
}

func (g *GetAccountByIdUseCase) Call(ctx context.Context, accountId uuid.UUID) (*entity.Account, error) {
	account, err := g.accountRepository.GetById(ctx, accountId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "GetAccountById"}).Error(err.Error())
		return nil, err
	}

	accountErrors := account.Validate()
	if accountErrors != nil {
		logrus.WithFields(logrus.Fields{"UseCase": "GetAccountById"}).Error(accountErrors.Error())
		return nil, g.MakeError(accountErrors.Error())
	}

	logrus.WithFields(logrus.Fields{"UseCase": "GetAccountById"}).Info("Success on get account: ", account.GetId().String())
	return account, nil
}

func (GetAccountByIdUseCase) MakeError(errorMessage string) error {
	return commonerrors.NewErrorInternalServer("GetAccountByIdUseCase", errorMessage)
}
