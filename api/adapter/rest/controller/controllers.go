package controller

import (
	"github.com/nelsonlpco/transactions/shared/dependencies"
)

type Controllers struct {
	AccountController     *AccountController
	TransactionController *TransactionController
}

func NewControllers(services *dependencies.Services) *Controllers {
	return &Controllers{
		AccountController:     NewAccountController(services.AccountService),
		TransactionController: NewTransactionController(services.TransactionService),
	}
}
