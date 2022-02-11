package datasource

import (
	"context"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type SqlOperationTypeDatasource struct {
	dbManger *db_manager.DBManager
}

func NewSqlOperationTypeDatasource(dbManger *db_manager.DBManager) *SqlOperationTypeDatasource {
	return &SqlOperationTypeDatasource{dbManger: dbManger}
}

func (s *SqlOperationTypeDatasource) Create(ctx context.Context, operationTypeModel *inframodel.OperationTypeModel) error {
	return nil
}

func (s *SqlOperationTypeDatasource) GetById(ctx context.Context, id []byte) (*inframodel.OperationTypeModel, error) {
	return &inframodel.OperationTypeModel{
		Id:          uuid.NewString(),
		Operation:   byte(valueobjects.Credit),
		Description: "PAGAMENTO",
	}, nil
}
