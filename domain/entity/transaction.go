package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
)

var ErrorTransactionAccountNotBeNil = errors.New("account is required, not be nil")
var ErrorTransactionOperationTypeNotBeNil = errors.New("OperationType is required, not be nil")

type Transaction struct {
	id            uuid.UUID
	amount        valueobjects.Money
	account       *Account
	operationType *OperationType
	eventDate     time.Time
}

func NewTransaction(id uuid.UUID, amount valueobjects.Money, account *Account, operationType *OperationType, eventDate time.Time) *Transaction {
	transaction := &Transaction{
		id:            id,
		account:       account,
		operationType: operationType,
		eventDate:     eventDate,
	}

	transaction.setAmount(amount)

	return transaction
}

func (t *Transaction) Validate() error {
	var errorMessages []string

	accountError := t.validateAccount()
	if accountError != "" {
		errorMessages = append(errorMessages, accountError)
	}

	operationTypeError := t.validateOperationType()
	if operationTypeError != "" {
		errorMessages = append(errorMessages, operationTypeError)
	}

	errorAmount := t.amount.Validate()
	if errorAmount != nil {
		errorMessages = append(errorMessages, errorAmount.Error())
	}

	if len(errorMessages) > 0 {
		return commonerrors.NewErrorInvalidEntity("Transaction", errorMessages)
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

func (t *Transaction) GetId() uuid.UUID {
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

func (t *Transaction) validateAccount() string {
	if t.account == nil {
		return ErrorTransactionAccountNotBeNil.Error()
	}

	err := t.account.Validate()
	if err != nil {
		return err.Error()
	}

	return ""
}

func (t *Transaction) validateOperationType() string {
	if t.operationType == nil {
		return ErrorTransactionOperationTypeNotBeNil.Error()
	}

	err := t.operationType.Validate()
	if err != nil {
		return err.Error()
	}

	return ""
}
