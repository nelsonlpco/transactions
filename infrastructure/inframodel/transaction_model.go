package inframodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type TransactionModel struct {
	Id            string
	Amount        float64
	EventDate     time.Time
	Account       *AccountModel
	OperationType *OperationTypeModel
}

func (t *TransactionModel) FromEntity(transaction *entity.Transaction) *TransactionModel {
	return &TransactionModel{
		Id:            transaction.GetId().String(),
		Amount:        float64(transaction.GetAmount()),
		EventDate:     transaction.GetEventDate(),
		Account:       new(AccountModel).FromEntity(transaction.GetAccount()),
		OperationType: new(OperationTypeModel).FromEntity(transaction.GetOperationType()),
	}
}

func (t *TransactionModel) ToEntity() *entity.Transaction {
	return entity.NewTransaction(
		uuid.MustParse(t.Id),
		valueobjects.NewMoney(t.Amount),
		t.Account.ToEntity(),
		t.OperationType.ToEntity(),
		t.EventDate,
	)
}
