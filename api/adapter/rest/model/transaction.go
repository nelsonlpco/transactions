package model

import "time"

type TransactionModel struct {
	Id          int       `json:"id"`
	Amount      float64   `json:"amount"`
	EventDate   time.Time `json:"eventDate"`
	AccountId   int       `json:"accountId"`
	OperationId int       `json:"operationId"`
}
