package commonerrors

import "fmt"

type ErrorNoContent struct {
	errorMessage string
}

func NewErrorNoContent(errorMessage string) *ErrorNoContent {
	return &ErrorNoContent{
		errorMessage: errorMessage,
	}
}

func (e *ErrorNoContent) Error() string {
	return fmt.Sprintf(`"%v"`, e.errorMessage)
}
