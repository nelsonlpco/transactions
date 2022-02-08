package entity

import (
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
)

type OperationType struct {
	Id          valueobjects.Id
	Description string
	Operation   valueobjects.Operation
}

func NewOperationType(id int, description string, operation byte) *OperationType {
	return &OperationType{
		Id:          valueobjects.NewId(id),
		Description: description,
		Operation:   valueobjects.NewOperation(operation),
	}
}

func (o *OperationType) Validate() []error {
	var validationErrors []error

	if !o.Id.IsValid() {
		validationErrors = append(validationErrors, domainerrors.NewErrorInvalidId("operationType"))
	}

	if !o.Operation.IsValid() {
		validationErrors = append(validationErrors, domainerrors.NewErrorInvalidOperation("operationType"))
	}

	if o.Description == "" {
		validationErrors = append(validationErrors, domainerrors.NewErrorInvalidDescription("operationType"))
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
