package domainerrors

import (
	"fmt"
)

type ErrorInvalidId struct {
	entityName string
}

func NewErrorInvalidId(entityName string) *ErrorInvalidId {
	return &ErrorInvalidId{
		entityName: entityName,
	}
}

func (e *ErrorInvalidId) Error() string {
	return fmt.Sprintf("%v: the id must be greater than 0", e.entityName)
}
