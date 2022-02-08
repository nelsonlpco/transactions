package domainerrors

import (
	"fmt"
)

type ErrorInvalidAmount struct {
	entityName string
}

func NewErrorInvalidAmount(entityName string) *ErrorInvalidAmount {
	return &ErrorInvalidAmount{
		entityName: entityName,
	}
}

func (e *ErrorInvalidAmount) Error() string {
	return fmt.Sprintf("%v: the amount must be different of zero ", e.entityName)
}
