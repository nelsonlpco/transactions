package controller

import "github.com/nelsonlpco/transactions/application/services"

type TransactionController struct {
	transactionService *services.TransactionService
}
