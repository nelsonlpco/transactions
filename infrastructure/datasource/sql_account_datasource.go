package datasource

import (
	"context"
	"fmt"

	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
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
		logrus.New().WithField("Datasource", "SqlAccountDataSource").Error(err)
		return fmt.Errorf(`"%v"`, err.Error())
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, accountModel.Id, accountModel.DocumentNumber)
	if err != nil {
		logrus.New().WithField("Datasource", "SqlAccountDataSource").Error(err)
		return fmt.Errorf(`"%v"`, err.Error())
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
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	var accountId []byte
	var documentNumber string
	if err := row.Scan(&accountId, &documentNumber); err != nil {
		logrus.New().WithField("Datasource", "SqlAccountDataSource").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	account := &inframodel.AccountModel{
		Id:             accountId,
		DocumentNumber: documentNumber,
	}

	return account, nil
}
