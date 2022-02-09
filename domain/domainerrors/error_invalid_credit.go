package domainerrors

import (
	"fmt"
)

type ErrorInvalidCredit struct {
	entityName string
}

func NewErrorInvalidCredit(entityName string) *ErrorInvalidCredit {
	return &ErrorInvalidCredit{
		entityName: entityName,
	}
}

func (e *ErrorInvalidCredit) Error() string {
	return fmt.Sprintf("%v: the credit must be positive", e.entityName)
}
