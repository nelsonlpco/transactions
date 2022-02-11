package dependencies

import (
	"github.com/nelsonlpco/transactions/domain/repository"
	repositoryImpl "github.com/nelsonlpco/transactions/infrastructure/repository"
)

type Repositories struct {
	AccountRepository       repository.AccountRepository
	OperationTypeRepository repository.OperationTypeRepository
	TransactionRepository   repository.TransactionRepository
}

func NewRepositories(datasources *Datasources) *Repositories {
	return &Repositories{
		AccountRepository:       repositoryImpl.NewAccountRepositoryImpl(datasources.AccountDatasource),
		OperationTypeRepository: repositoryImpl.NewOperationTypeRepositoryImpl(datasources.OperatonTypeDatasource),
		TransactionRepository:   repositoryImpl.NewTransactionRepositoryImpl(datasources.TransactionDatasource),
	}
}
