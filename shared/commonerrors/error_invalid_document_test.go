package commonerrors_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_invalid_document_error(t *testing.T) {
	err := commonerrors.NewErrorInvalidDocument("092340898")

	require.NotNil(t, err)
}

func Test_should_be_create_an_invalid_document_text(t *testing.T) {
	err := commonerrors.NewErrorInvalidDocument("092340898")

	require.Equal(t, `"092340898 is not a valid document"`, err.Error())
}
