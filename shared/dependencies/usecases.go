package dependencies

import (
	"github.com/nelsonlpco/transactions/domain/usecases"
)

type UseCases struct {
	CreateAccountUseCase        *usecases.CreateAccountUseCase
	GetAccountByIdUseCase       *usecases.GetAccountByIdUseCase
	CreateTransactionUseCase    *usecases.CreateTransactionUseCase
	GetOperationTypeByIdUseCase *usecases.GetOperationTypeByIdUseCase
}

func NewUseCases(repositories *Repositories) *UseCases {
	return &UseCases{
		CreateAccountUseCase:        usecases.NewCreateAccountUseCase(repositories.AccountRepository),
		GetAccountByIdUseCase:       usecases.NewGetAccountByIdUseCase(repositories.AccountRepository),
		CreateTransactionUseCase:    usecases.NewCreateTransactionUseCase(repositories.TransactionRepository),
		GetOperationTypeByIdUseCase: usecases.NewGetOperationTypeByIdUseCase(repositories.OperationTypeRepository),
	}
}
