package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/sirupsen/logrus"
)

type SqlOperationTypeDatasource struct {
	dbManger *db_manager.DBManager
}

func NewSqlOperationTypeDatasource(dbManger *db_manager.DBManager) *SqlOperationTypeDatasource {
	return &SqlOperationTypeDatasource{dbManger: dbManger}
}

func (s *SqlOperationTypeDatasource) Create(ctx context.Context, operationTypeModel *inframodel.OperationTypeModel) error {
	ctx, cancel := context.WithTimeout(ctx, s.dbManger.GetTTL())
	defer cancel()

	query := "INSERT INTO operation_type(id, description, operation) VALUES(?,?,?);"
	stmt, err := s.dbManger.GetDB().PrepareContext(ctx, query)
	if err != nil {
		logrus.New().WithContext(ctx).WithField("SqlOperationTypeDatasource", "Create").Error(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, operationTypeModel.Id, operationTypeModel.Description, operationTypeModel.Operation)
	if err != nil {
		logrus.New().WithContext(ctx).WithField("SqlOperationTypeDatasource", "Create").Error(err)
		return err
	}

	return nil
}

func (s *SqlOperationTypeDatasource) GetById(ctx context.Context, id []byte) (*inframodel.OperationTypeModel, error) {
	ctx, cancel := context.WithTimeout(ctx, s.dbManger.GetTTL())
	defer cancel()

	query := "SELECT * FROM operation_type WHERE id = ?"
	stmt, err := s.dbManger.GetDB().PrepareContext(ctx, query)
	if err != nil {
		logrus.New().WithContext(ctx).WithField("SqlOperationTypeDatasource", "GetById").Error(err)
		return nil, err
	}
	defer stmt.Close()

	var operationTypeModel inframodel.OperationTypeModel

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&operationTypeModel.Id, &operationTypeModel.Description, &operationTypeModel.Operation)
	if err != nil {
		logrus.New().WithContext(ctx).WithField("SqlOperationTypeDatasource", "GetById").Error(err)
		return nil, commonerrors.NewErrorNoContent("Operation Type not found")
	}

	return &operationTypeModel, nil
}
