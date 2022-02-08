package entity

import (
	"github.com/klassmann/cpfcnpj"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type Account struct {
	Id             valueobjects.Id
	DocumentNumber cpfcnpj.CPF
}

func NewAccount(id int, documentNumber string) *Account {
	return &Account{
		Id:             valueobjects.NewId(id),
		DocumentNumber: cpfcnpj.NewCPF(documentNumber),
	}
}

func (a *Account) Validate() []error {
	var validationerrors []error

	if !a.DocumentNumber.IsValid() {
		validationerrors = append(validationerrors, domainerrors.NewErrorInvalidDocument("account", a.DocumentNumber.String()))
	}

	if !a.Id.IsValid() {
		validationerrors = append(validationerrors, domainerrors.NewErrorInvalidId("account"))
	}

	if len(validationerrors) > 0 {
		return validationerrors
	}

	return nil
}
