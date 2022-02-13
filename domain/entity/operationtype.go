package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
)

var ErrorOperationTypeDescriptionRequired = errors.New(`"Description is required, not be empty"`)

type OperationType struct {
	id          uuid.UUID
	description string
	operation   valueobjects.Operation
}

func NewOperationType(id uuid.UUID, description string, operation valueobjects.Operation) *OperationType {
	return &OperationType{
		id:          id,
		description: description,
		operation:   operation,
	}
}

func (o *OperationType) Validate() error {
	var messageErrors []string

	operationError := o.operation.Validate()
	if operationError != nil {
		messageErrors = append(messageErrors, operationError.Error())
	}

	if o.description == "" {
		messageErrors = append(messageErrors, ErrorOperationTypeDescriptionRequired.Error())
	}

	if len(messageErrors) > 0 {
		return commonerrors.NewErrorInvalidEntity("OperationType", messageErrors)
	}

	return nil
}

func (o *OperationType) GetId() uuid.UUID {
	return o.id
}

func (o *OperationType) GetDescription() string {
	return o.description
}

func (o *OperationType) GetOperation() valueobjects.Operation {
	return o.operation
}
