package model

import (
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
)

type AccountModel struct {
	Id             string `json:"id"`
	DocumentNumber string `json:"documentNumber"`
}

func (a *AccountModel) FromEntity(account *entity.Account) *AccountModel {
	a.Id = account.GetId().String()
	a.DocumentNumber = account.GetDocumentNumber()
	return a
}

func (a *AccountModel) ToEntity() *entity.Account {
	return entity.NewAccount(uuid.MustParse(a.Id), a.DocumentNumber)
}
