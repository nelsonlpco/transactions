package domainerrors

import (
	"errors"
	"fmt"
)

func ErrorsToError(inputErrors []error) error {
	var errorMessage string

	for _, err := range inputErrors {
		errorMessage += fmt.Sprintf("%v; ", err.Error())
	}

	return errors.New(errorMessage)
}
