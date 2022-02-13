package inframodel

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/sirupsen/logrus"
)

type TransactionModel struct {
	Id            []byte
	Amount        float64
	EventDate     time.Time
	Account       *AccountModel
	OperationType *OperationTypeModel
}

func (t *TransactionModel) FromEntity(transaction *entity.Transaction) (*TransactionModel, error) {
	account, err := new(AccountModel).FromEntity(transaction.GetAccount())
	if err != nil {
		logrus.WithField("TransactionModel", "FromEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	OperationType, err := new(OperationTypeModel).FromEntity(transaction.GetOperationType())
	if err != nil {
		logrus.WithField("TransactionModel", "FromEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	binaryTransactionTypeId, err := transaction.GetId().MarshalBinary()
	if err != nil {
		logrus.WithField("TransactionModel", "FromEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	return &TransactionModel{
		Id:            binaryTransactionTypeId,
		Amount:        float64(transaction.GetAmount()),
		EventDate:     transaction.GetEventDate(),
		Account:       account,
		OperationType: OperationType,
	}, nil
}

func (t *TransactionModel) ToEntity() (*entity.Transaction, error) {
	account, err := t.Account.ToEntity()
	if err != nil {
		logrus.WithField("TransactionModel", "ToEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	operationType, err := t.OperationType.ToEntity()
	if err != nil {
		logrus.WithField("TransactionModel", "ToEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	transactionId := uuid.New()
	err = transactionId.UnmarshalBinary(t.Id)
	if err != nil {
		logrus.WithField("TransactionModel", "ToEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	return entity.NewTransaction(
		transactionId,
		valueobjects.NewMoney(t.Amount),
		account,
		operationType,
		t.EventDate,
	), nil
}
