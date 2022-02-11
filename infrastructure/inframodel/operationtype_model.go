package inframodel

import (
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type OperationTypeModel struct {
	Id          string
	Description string
	Operation   byte
}

func (a *OperationTypeModel) FromEntity(operationType *entity.OperationType) *OperationTypeModel {
	a.Id = operationType.GetId().String()
	a.Description = operationType.GetDescription()
	a.Operation = byte(operationType.GetOperation())

	return a
}

func (a *OperationTypeModel) ToEntity() *entity.OperationType {
	return entity.NewOperationType(uuid.MustParse(a.Id), a.Description, valueobjects.Operation(a.Operation))
}
