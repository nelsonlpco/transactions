package domainerrors

import "fmt"

type ErrorInvalidDescription struct {
	entityName string
}

func NewErrorInvalidDescription(entityName string) *ErrorInvalidDescription {
	return &ErrorInvalidDescription{
		entityName: entityName,
	}
}

func (e *ErrorInvalidDescription) Error() string {
	return fmt.Sprintf("%v: the description is not be empty", e.entityName)
}
