package domainerrors

import (
	"fmt"
)

type ErrorInvalidDebit struct {
	entityName string
}

func NewErrorInvalidDebit(entityName string) *ErrorInvalidDebit {
	return &ErrorInvalidDebit{
		entityName: entityName,
	}
}

func (e *ErrorInvalidDebit) Error() string {
	return fmt.Sprintf("%v: the debit must be negative", e.entityName)
}
