package repository

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
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
	accountModel := &inframodel.AccountModel{
		Id:             int(account.GetId()),
		DocumentNumber: account.GetDocumentNumber(),
	}

	err := a.accountDatasource.Create(ctx, accountModel)
	if err != nil {
		return fmt.Errorf("AccountRepositoryImpl: %v", err)
	}

	return nil
}

func (a *AccountRepositoryImpl) GetById(ctx context.Context, id valueobjects.Id) (*entity.Account, error) {
	accountModel, err := a.accountDatasource.GetById(ctx, int(id))

	if err != nil {
		return nil, fmt.Errorf("AccountRepositoryImpl: %v", err)
	}

	accountEntity := entity.NewAccount(valueobjects.Id(accountModel.Id), accountModel.DocumentNumber)

	return accountEntity, nil
}
