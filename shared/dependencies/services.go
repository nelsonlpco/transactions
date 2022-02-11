package dependencies

import "github.com/nelsonlpco/transactions/application/services"

type Services struct {
	AccountService     *services.AccountService
	TransactionService *services.TransactionService
}

func NewServices(useCases *UseCases) *Services {
	return &Services{
		AccountService: services.NewAccountService(useCases.GetAccountByIdUseCase, useCases.CreateAccountUseCase),
		TransactionService: services.NewTransactionService(
			useCases.GetAccountByIdUseCase,
			useCases.GetOperationTypeByIdUseCase,
			useCases.CreateTransactionUseCase),
	}
}
