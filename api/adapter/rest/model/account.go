package model

import (
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type AccountModel struct {
	Id             int    `json:"id"`
	DocumentNumber string `json:"documentNumber"`
}

func (a *AccountModel) FromEntity(account *entity.Account) *AccountModel {
	a.Id = int(account.GetId())
	a.DocumentNumber = account.GetDocumentNumber()
	return a
}

func (a *AccountModel) ToEntity() *entity.Account {
	return entity.NewAccount(valueobjects.Id(a.Id), a.DocumentNumber)
}
