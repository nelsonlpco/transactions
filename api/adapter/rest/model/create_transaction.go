package model

type CreateTransactionModel struct {
	Amount          float64 `json:"amount"`
	AccountId       string  `json:"accountId"`
	OperationTypeId string  `json:"operationTypeId"`
}
