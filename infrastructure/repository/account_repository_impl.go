package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/sirupsen/logrus"
)

type AccountRepositoryImpl struct {
	accountDatasource datasource.AccountDatasource
}

func NewAccountRepositoryImpl(accountDatasource datasource.AccountDatasource) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		accountDatasource: accountDatasource,
	}
}

func (a *AccountRepositoryImpl) Create(ctx context.Context, account *entity.Account) error {
	logrus.Info(account.GetId().String())
	accountModel := new(inframodel.AccountModel).FromEntity(account)

	err := a.accountDatasource.Create(ctx, accountModel)
	if err != nil {
		return a.MakeError(err.Error())
	}

	return nil
}

func (a *AccountRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	byteId, err := id.MarshalBinary()
	if err != nil {
		return nil, a.MakeError(err.Error())
	}

	accountModel, err := a.accountDatasource.GetById(ctx, byteId)
	if err != nil {
		return nil, a.MakeError(err.Error())
	}

	accountEntity := accountModel.ToEntity()

	return accountEntity, nil
}

func (AccountRepositoryImpl) MakeError(errorMessage string) error {
	return domainerrors.NewErrorInternalServer("AccountRepositoryImpl", errorMessage)
}
