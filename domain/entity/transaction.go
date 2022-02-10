package entity

import (
	"errors"
	"time"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type Transaction struct {
	id            valueobjects.Id
	amount        valueobjects.Money
	account       *Account
	operationType *OperationType
	eventDate     time.Time
}

func NewTransaction(id valueobjects.Id, amount valueobjects.Money, account *Account, operationType *OperationType, eventDate time.Time) *Transaction {
	transaction := &Transaction{
		id:            id,
		account:       account,
		operationType: operationType,
		eventDate:     eventDate,
	}

	transaction.setAmount(amount)

	return transaction
}

func (t *Transaction) Validate() []error {
	var validationError []error

	if !t.id.IsValid() {
		validationError = append(validationError, domainerrors.NewErrorInvalidId("transaction"))
	}

	errorAmount := t.validateAmount()
	if errorAmount != nil {
		validationError = append(validationError, errorAmount)
	}

	if t.account == nil {
		validationError = append(validationError, errors.New("account is required"))
	} else {
		accountErrors := t.account.Validate()
		validationError = append(validationError, accountErrors...)
	}

	if t.operationType == nil {
		validationError = append(validationError, errors.New("operation type is required"))
	} else {
		operationTypeErrors := t.operationType.Validate()
		validationError = append(validationError, operationTypeErrors...)
	}

	if len(validationError) > 0 {
		return validationError
	}

	return nil
}

func (t *Transaction) setAmount(amount valueobjects.Money) {
	if t.operationType == nil {
		t.amount = amount
		return
	}

	if t.operationType.operation.IsCredit() && amount < 0 {
		amount *= -1
	}
	if t.operationType.operation.IsDebit() && amount > 0 {
		amount *= -1
	}

	t.amount = amount
}

func (t *Transaction) validateAmount() error {
	if t.amount == 0 {
		return domainerrors.NewErrorInvalidAmount("transaction")
	}

	return nil
}

func (t *Transaction) GetId() valueobjects.Id {
	return t.id
}

func (t *Transaction) GetAmount() valueobjects.Money {
	return t.amount
}

func (t *Transaction) GetEventDate() time.Time {
	return t.eventDate
}

func (t *Transaction) GetAccount() *Account {
	return t.account
}

func (t *Transaction) GetOperationType() *OperationType {
	return t.operationType
}
