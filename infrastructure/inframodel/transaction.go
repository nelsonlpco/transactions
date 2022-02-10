package inframodel

import "time"

type TransactionModel struct {
	Id            int
	Amount        float64
	EventDate     time.Time
	Account       *AccountModel
	OperationType *OperationTypeModel
}
