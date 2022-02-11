package datasource

import (
	"context"
	"errors"

	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type SqlTransactionDatasource struct {
	dbManger *db_manager.DBManager
}

func NewSqlTransactionDatasource(dbManger *db_manager.DBManager) *SqlTransactionDatasource {
	return &SqlTransactionDatasource{dbManger: dbManger}
}

func (s *SqlTransactionDatasource) Create(ctx context.Context, transaction *inframodel.TransactionModel) error {
	return errors.New(`"eita"`)
}
