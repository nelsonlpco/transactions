package domainerrors

import "fmt"

type ErrorInvalidDocument struct {
	documentNumber string
}

func NewErrorInvalidDocument(documentNumber string) *ErrorInvalidDocument {
	return &ErrorInvalidDocument{
		documentNumber: documentNumber,
	}
}

func (e *ErrorInvalidDocument) Error() string {
	return fmt.Sprintf(`"%v is not a valid document"`, e.documentNumber)
}
