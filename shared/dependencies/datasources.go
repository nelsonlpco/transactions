package dependencies

import (
	"github.com/nelsonlpco/transactions/infrastructure/datasource"
	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
)

type Datasources struct {
	AccountDatasource      datasource.AccountDatasource
	OperatonTypeDatasource datasource.OperationTypeDatasource
	TransactionDatasource  datasource.TransactionDatasource
}

func NewSqlDatasources(dbManager *db_manager.DBManager) *Datasources {
	return &Datasources{
		AccountDatasource:      datasource.NewSqlAccountDatasource(dbManager),
		OperatonTypeDatasource: datasource.NewSqlOperationTypeDatasource(dbManager),
		TransactionDatasource:  datasource.NewSqlTransactionDatasource(dbManager),
	}
}
