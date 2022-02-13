package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
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
	accountModel, err := new(inframodel.AccountModel).FromEntity(account)
	if err != nil {
		logrus.WithField("accountRepository", "Crate").Error(err)
		return a.MakeError(err.Error())
	}

	err = a.accountDatasource.Create(ctx, accountModel)
	if err != nil {
		logrus.WithField("accountRepository", "Crate").Error(err)
		return err
	}

	return nil
}

func (a *AccountRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	byteId, err := id.MarshalBinary()
	if err != nil {
		logrus.WithField("accountRepository", "GetById").Error(err)
		return nil, a.MakeError(err.Error())
	}

	accountModel, err := a.accountDatasource.GetById(ctx, byteId)
	if err != nil {
		logrus.WithField("accountRepository", "GetById").Error(err)
		return nil, err
	}

	accountEntity, err := accountModel.ToEntity()
	if err != nil {
		logrus.WithField("accountRepository", "GetById").Error(err)
		return nil, a.MakeError(err.Error())
	}

	return accountEntity, nil
}

func (a *AccountRepositoryImpl) GetByDocumentNumber(ctx context.Context, documentNumber string) (*entity.Account, error) {
	accountModel, err := a.accountDatasource.GetByDocumentNumber(ctx, documentNumber)
	if err != nil {
		logrus.WithField("accountRepository", "GetByDocumentNumber").Error(err)
		return nil, err
	}

	accountEntity, err := accountModel.ToEntity()
	if err != nil {
		logrus.WithField("accountRepository", "GetByDocumentNumber").Error(err)
		return nil, a.MakeError(err.Error())
	}

	return accountEntity, nil
}

func (AccountRepositoryImpl) MakeError(errorMessage string) error {
	return commonerrors.NewErrorInternalServer("AccountRepositoryImpl", errorMessage)
}
