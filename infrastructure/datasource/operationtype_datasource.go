package datasource

import (
	"context"

	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
)

type OperationTypeDatasource interface {
	Create(ctx context.Context, operationType *inframodel.OperationTypeModel) error
	GetById(ctx context.Context, id []byte) (*inframodel.OperationTypeModel, error)
}
