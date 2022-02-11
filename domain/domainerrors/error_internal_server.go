package domainerrors

import "fmt"

type ErrorInternalServer struct {
	operationName string
	errorMessage  string
}

func NewErrorInternalServer(operationName, errorMessage string) *ErrorInternalServer {
	return &ErrorInternalServer{
		operationName: operationName,
		errorMessage:  errorMessage,
	}
}

func (e *ErrorInternalServer) Error() string {
	return fmt.Sprintf(`{"%v": %v}`, e.operationName, e.errorMessage)
}
