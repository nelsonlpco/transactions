package dependencies

import "github.com/nelsonlpco/transactions/api/adapter/rest/controller"

type Controllers struct {
	AccountController     *controller.AccountController
	TransactionController *controller.TransactionController
}

func NewControllers(services *Services) *Controllers {
	return &Controllers{
		AccountController:     controller.NewAccountController(services.AccountService),
		TransactionController: controller.NewTransactionController(services.TransactionService),
	}
}
