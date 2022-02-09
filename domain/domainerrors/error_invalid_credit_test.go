package domainerrors_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_invalid_credit_error(t *testing.T) {
	err := domainerrors.NewErrorInvalidCredit("teste")

	require.NotNil(t, err)

}