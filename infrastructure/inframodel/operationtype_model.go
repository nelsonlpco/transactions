package inframodel

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/sirupsen/logrus"
)

type OperationTypeModel struct {
	Id          []byte
	Description string
	Operation   byte
}

func (a *OperationTypeModel) FromEntity(operationType *entity.OperationType) (*OperationTypeModel, error) {
	binaryId, err := operationType.GetId().MarshalBinary()
	if err != nil {
		logrus.New().WithField("OperationTypeModel", "FromEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	a.Id = binaryId
	a.Description = operationType.GetDescription()
	a.Operation = byte(operationType.GetOperation())

	return a, nil
}

func (a *OperationTypeModel) ToEntity() (*entity.OperationType, error) {
	id := uuid.New()
	err := id.UnmarshalBinary(a.Id)
	if err != nil {
		logrus.New().WithField("OperationTypeModel", "ToEntity").Error(err)
		return nil, fmt.Errorf(`"%v"`, err.Error())
	}

	return entity.NewOperationType(id, a.Description, valueobjects.Operation(a.Operation)), nil
}
