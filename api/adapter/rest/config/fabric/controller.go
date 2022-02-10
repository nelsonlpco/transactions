package fabric

import (
	"github.com/nelsonlpco/transactions/api/adapter/rest/controller"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
)

func MakeAccountController() *controller.AccountController {
	datasource := datasource.NewSqlAccountDatasource()
	accountRepository := repository.NewAccountRepositoryImpl(datasource)
	getAccountByIdUseCase := usecases.NewGetAccountById(accountRepository)
	createAccount := usecases.NewCreateAccount(accountRepository)

	accountService := services.NewAccountService(getAccountByIdUseCase, createAccount)

	return controller.NewAccountController(accountService)
}
