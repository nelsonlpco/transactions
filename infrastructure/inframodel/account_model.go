package inframodel

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/sirupsen/logrus"
)

type AccountModel struct {
	Id             []byte
	DocumentNumber string
}

func (a *AccountModel) FromEntity(account *entity.Account) (*AccountModel, error) {
	binaryAccountId, err := account.GetId().MarshalBinary()
	if err != nil {
		logrus.New().WithField("AccountModel", "FromEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	a.Id = binaryAccountId
	a.DocumentNumber = account.GetDocumentNumber()

	return a, nil
}

func (a *AccountModel) ToEntity() (*entity.Account, error) {
	accountId := uuid.New()
	err := accountId.UnmarshalBinary(a.Id)
	if err != nil {
		logrus.New().WithField("AccountModel", "ToEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}
	return entity.NewAccount(accountId, a.DocumentNumber), nil
}
