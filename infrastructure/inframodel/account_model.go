package inframodel

import (
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
)

type AccountModel struct {
	Id             []byte
	DocumentNumber string
}

func (a *AccountModel) FromEntity(account *entity.Account) *AccountModel {
	id, _ := account.GetId().MarshalBinary()

	a.Id = id
	a.DocumentNumber = account.GetDocumentNumber()

	return a
}

func (a *AccountModel) ToEntity() *entity.Account {
	id := uuid.New()
	id.UnmarshalBinary(a.Id)
	return entity.NewAccount(id, a.DocumentNumber)
}
