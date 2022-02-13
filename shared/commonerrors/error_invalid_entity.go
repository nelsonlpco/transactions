package commonerrors

import (
	"fmt"
	"strings"
)

type ErrorInvalidEntity struct {
	entityName    string
	errorMessages []string
}

func NewErrorInvalidEntity(entityName string, errorMessages []string) *ErrorInvalidEntity {
	return &ErrorInvalidEntity{
		entityName:    entityName,
		errorMessages: errorMessages,
	}
}

func (e *ErrorInvalidEntity) Error() string {
	return fmt.Sprintf(`{"%v": [%v]}`, e.entityName, strings.Join(e.errorMessages, ","))
}
