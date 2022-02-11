package entity

import (
	"github.com/google/uuid"
	"github.com/klassmann/cpfcnpj"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
)

type Account struct {
	id             uuid.UUID
	documentNumber cpfcnpj.CPF
}

func NewAccount(id uuid.UUID, documentNumber string) *Account {
	return &Account{
		id:             id,
		documentNumber: cpfcnpj.NewCPF(documentNumber),
	}
}

func (a *Account) Validate() error {
	if !a.documentNumber.IsValid() {
		err := domainerrors.NewErrorInvalidDocument(a.GetDocumentNumber())
		messageErrors := []string{err.Error()}
		return domainerrors.NewErrorInvalidEntity("Account", messageErrors)
	}

	return nil
}

func (a Account) GetId() uuid.UUID {
	return a.id
}

func (a Account) GetDocumentNumber() string {
	return string(a.documentNumber)
}
