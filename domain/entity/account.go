package entity

import (
	"github.com/klassmann/cpfcnpj"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type Account struct {
	id             valueobjects.Id
	documentNumber cpfcnpj.CPF
}

func NewAccount(id valueobjects.Id, documentNumber string) *Account {
	return &Account{
		id:             id,
		documentNumber: cpfcnpj.NewCPF(documentNumber),
	}
}

func (a *Account) Validate() []error {
	var validationerrors []error

	if !a.documentNumber.IsValid() {
		validationerrors = append(validationerrors, domainerrors.NewErrorInvalidDocument("account", a.GetDocumentNumber()))
	}

	if !a.id.IsValid() {
		validationerrors = append(validationerrors, domainerrors.NewErrorInvalidId("account"))
	}

	if len(validationerrors) > 0 {
		return validationerrors
	}

	return nil
}

func (a Account) GetId() valueobjects.Id {
	return a.id
}

func (a Account) GetDocumentNumber() string {
	return string(a.documentNumber)
}
