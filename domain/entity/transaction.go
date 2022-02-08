package entity

import (
	"time"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type Transaction struct {
	Id            valueobjects.Id
	Amount        float64
	Account       *Account
	OperationType *OperationType
	EventDate     time.Time
}

func NewTransaction(id int, amount float64, account *Account, operationType *OperationType, eventDate time.Time) *Transaction {
	return &Transaction{
		Id:            valueobjects.NewId(id),
		Amount:        amount,
		Account:       account,
		OperationType: operationType,
		EventDate:     eventDate,
	}
}

func (t *Transaction) Validate() []error {
	var validationError []error

	if !t.Id.IsValid() {
		validationError = append(validationError, domainerrors.NewErrorInvalidId("transaction"))
	}

	if t.Amount == 0 {
		validationError = append(validationError, domainerrors.NewErrorInvalidAmount("transaction"))
	}

	accountErrors := t.Account.Validate()
	validationError = append(validationError, accountErrors...)

	operationTypeErrors := t.OperationType.Validate()
	validationError = append(validationError, operationTypeErrors...)

	if len(validationError) > 0 {
		return validationError
	}

	return nil
}
