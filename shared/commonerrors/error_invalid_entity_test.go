package commonerrors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_invalid_entity_error(t *testing.T) {
	errorMessages := []string{`"error 1"`, `"error 2"`}
	err := commonerrors.NewErrorInvalidEntity("account", errorMessages)
	require.NotNil(t, err)
}

func Test_should_be_create_an_invalid_entity_text(t *testing.T) {
	errorMessages := []string{`"error 1"`, `"error 2"`}
	err := commonerrors.NewErrorInvalidEntity("account", errorMessages)

	expectedError := fmt.Sprintf(`{"%v": [%v]}`, "account", strings.Join(errorMessages, ","))

	require.Equal(t, expectedError, err.Error())
}
