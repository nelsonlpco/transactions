package entity

import (
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type OperationType struct {
	id          valueobjects.Id
	description string
	operation   valueobjects.Operation
}

func NewOperationType(id valueobjects.Id, description string, operation valueobjects.Operation) *OperationType {
	return &OperationType{
		id:          id,
		description: description,
		operation:   operation,
	}
}

func (o *OperationType) Validate() []error {
	var validationErrors []error

	if !o.id.IsValid() {
		validationErrors = append(validationErrors, domainerrors.NewErrorInvalidId("operationType"))
	}

	if !o.operation.IsValid() {
		validationErrors = append(validationErrors, domainerrors.NewErrorInvalidOperation("operationType"))
	}

	if o.description == "" {
		validationErrors = append(validationErrors, domainerrors.NewErrorInvalidDescription("operationType"))
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func (o *OperationType) GetId() valueobjects.Id {
	return o.id
}

func (o *OperationType) GetDescription() string {
	return o.description
}

func (o *OperationType) GetOperation() valueobjects.Operation {
	return o.operation
}
