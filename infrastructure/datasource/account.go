package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type AccountDatasource interface {
	Create(ctx context.Context, accountModel *inframodel.AccountModel) error
	GetById(ctx context.Context, id int) (*inframodel.AccountModel, error)
}
