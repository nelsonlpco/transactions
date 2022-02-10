package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type SqlAccountDatasource struct {
}

func NewSqlAccountDatasource() *SqlAccountDatasource {
	return &SqlAccountDatasource{}
}

func (s *SqlAccountDatasource) Create(ctx context.Context, accountModel *inframodel.AccountModel) error {
	return nil
}

func (s *SqlAccountDatasource) GetById(ctx context.Context, id int) (*inframodel.AccountModel, error) {
	return &inframodel.AccountModel{
		Id:             1,
		DocumentNumber: "07214684977",
	}, nil
}
