package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/sirupsen/logrus"
)

type SqlTransactionDatasource struct {
	dbManger *db_manager.DBManager
}

func NewSqlTransactionDatasource(dbManger *db_manager.DBManager) *SqlTransactionDatasource {
	return &SqlTransactionDatasource{dbManger: dbManger}
}

func (s *SqlTransactionDatasource) Create(ctx context.Context, transaction *inframodel.TransactionModel) error {
	ctx, cancel := context.WithTimeout(ctx, s.dbManger.GetTTL())
	defer cancel()

	query := "INSERT INTO transactions(id, account_id, operation_type_id, amount, event_date) VALUES(?,?,?,?,?)"
	stmt, err := s.dbManger.GetDB().PrepareContext(ctx, query)
	if err != nil {
		logrus.New().WithContext(ctx).WithField("SqlTransactionDatasource", "Create").Error(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		transaction.Id,
		transaction.Account.Id,
		transaction.OperationType.Id,
		transaction.Amount,
		transaction.EventDate,
	)

	if err != nil {
		logrus.New().WithContext(ctx).WithField("SqlTransactionDatasource", "Create").Error(err)
		return err
	}

	return nil
}
