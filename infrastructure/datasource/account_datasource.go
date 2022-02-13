package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type AccountDatasource interface {
	Create(ctx context.Context, accountModel *inframodel.AccountModel) error
	GetById(ctx context.Context, id []byte) (*inframodel.AccountModel, error)
	GetByDocumentNumber(ctx context.Context, documentNumber string) (*inframodel.AccountModel, error)
}
