package valueobjects_test

import (
	"errors"
	"testing"

	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_a_credit_operation(t *testing.T) {
	operation := valueobjects.NewOperation(valueobjects.Credit)

	require.NotNil(t, operation)
	require.Equal(t, valueobjects.Credit, operation)
	require.False(t, operation.IsDebit())
	require.True(t, operation.IsCredit())
	require.Nil(t, operation.Validate())
}

func Test_should_be_create_a_debit_operation(t *testing.T) {
	operation := valueobjects.NewOperation(valueobjects.Debit)

	require.True(t, operation.IsDebit())
	require.False(t, operation.IsCredit())
	require.Equal(t, valueobjects.Debit, operation)
	require.Nil(t, operation.Validate())
}

func Test_should_be_create_an_invalid_operation(t *testing.T) {
	operation := valueobjects.NewOperation(2)
	err := operation.Validate()

	require.True(t, errors.As(err, &valueobjects.ErrorInvalidOperation))
	require.False(t, operation.IsCredit())
	require.False(t, operation.IsDebit())

}
