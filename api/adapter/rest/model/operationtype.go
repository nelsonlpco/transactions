package model

import (
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type OperationTypeModel struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Operation   byte   `json:"operation"`
}

func (o *OperationTypeModel) ToEntity() *entity.OperationType {
	return entity.NewOperationType(
		uuid.MustParse(o.Id),
		o.Description,
		valueobjects.Operation(o.Operation),
	)
}

func (o *OperationTypeModel) FromEntity(operationType *entity.OperationType) *OperationTypeModel {
	o.Id = operationType.GetId().String()
	o.Description = operationType.GetDescription()
	o.Operation = byte(operationType.GetOperation())

	return o
}
