package domainerrors

import "fmt"

type ErrorInvalidDocument struct {
	entityName     string
	documentNumber string
}

func NewErrorInvalidDocument(entityName, documentNumber string) *ErrorInvalidDocument {
	return &ErrorInvalidDocument{
		documentNumber: documentNumber,
		entityName:     entityName,
	}
}

func (e *ErrorInvalidDocument) Error() string {
	return fmt.Sprintf("%v: %v is not a valid document.", e.entityName, e.documentNumber)
}
