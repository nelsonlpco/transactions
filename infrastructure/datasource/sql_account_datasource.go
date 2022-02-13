package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/sirupsen/logrus"
)

type SqlAccountDatasource struct {
	dbManger *db_manager.DBManager
}

func NewSqlAccountDatasource(dbManager *db_manager.DBManager) *SqlAccountDatasource {
	return &SqlAccountDatasource{dbManger: dbManager}
}

func (s *SqlAccountDatasource) Create(ctx context.Context, accountModel *inframodel.AccountModel) error {
	ctx, cancel := context.WithTimeout(ctx, s.dbManger.GetTTL())
	defer cancel()

	query := "INSERT INTO account(id, document_number) values(?,?)"
	logrus.Trace(query)

	stmt, err := s.dbManger.GetDB().PrepareContext(ctx, query)
	if err != nil {
		logrus.New().WithField("SqlAccountDataSource", "Create").Error(err)
		return commonerrors.NewErrorInternalServer("SQL", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, accountModel.Id, accountModel.DocumentNumber)
	if err != nil {
		logrus.New().WithField("Datasource", "SqlAccountDataSource").Error(err)
		return commonerrors.HandleSqlError(err)
	}
	return nil
}

func (s *SqlAccountDatasource) GetById(ctx context.Context, id []byte) (*inframodel.AccountModel, error) {
	ctx, cancel := context.WithTimeout(ctx, s.dbManger.GetTTL())
	defer cancel()

	query := "SELECT * FROM account WHERE id = ?"
	stmt, err := s.dbManger.GetDB().PrepareContext(ctx, query)
	if err != nil {
		logrus.New().WithField("Datasource", "SqlAccountDataSource").Error(err)
		return nil, commonerrors.NewErrorInternalServer("sql", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	var accountModel inframodel.AccountModel

	err = row.Scan(&accountModel.Id, &accountModel.DocumentNumber)
	if err != nil {
		logrus.New().WithField("SqlAccountDataSource", "GetById").Warn(err)
		return nil, commonerrors.NewErrorNoContent("Account not found")
	}

	return &accountModel, nil
}

func (s *SqlAccountDatasource) GetByDocumentNumber(ctx context.Context, documentNumber string) (*inframodel.AccountModel, error) {
	ctx, cancel := context.WithTimeout(ctx, s.dbManger.GetTTL())
	defer cancel()

	query := "SELECT * FROM account WHERE document_number = ?"
	stmt, err := s.dbManger.GetDB().PrepareContext(ctx, query)
	if err != nil {
		logrus.New().WithField("Datasource", "SqlAccountDataSource").Error(err)
		return nil, commonerrors.NewErrorInternalServer("sql", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, documentNumber)
	var accountModel inframodel.AccountModel

	err = row.Scan(&accountModel.Id, &accountModel.DocumentNumber)
	if err != nil {
		logrus.New().WithField("SqlAccountDataSource", "GetByDocumentNumber").Error(err)
		return nil, commonerrors.NewErrorNoContent("Account not found")
	}

	return &accountModel, nil
}
