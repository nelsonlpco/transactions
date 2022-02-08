package domainerrors

import (
	"fmt"
)

type ErrorInvalidOperation struct {
	entityName string
}

func NewErrorInvalidOperation(entityName string) *ErrorInvalidOperation {
	return &ErrorInvalidOperation{
		entityName: entityName,
	}
}

func (e *ErrorInvalidOperation) Error() string {
	return fmt.Sprintf("%v: operation must be Debit(0) or Credit(1)", e.entityName)
}
