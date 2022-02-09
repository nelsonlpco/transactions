package valueobjects_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_operation(t *testing.T) {
	operation := valueobjects.NewOperation(valueobjects.Credit)

	require.NotNil(t, operation)
	require.Equal(t, valueobjects.Credit, operation)
	require.False(t, operation.IsDebit())
	require.True(t, operation.IsCredit())
}

func Test_should_be_create_an_valid_operation(t *testing.T) {
	operation := valueobjects.NewOperation(valueobjects.Debit)

	require.True(t, operation.IsValid())
	require.True(t, operation.IsDebit())
	require.False(t, operation.IsCredit())
	require.Equal(t, valueobjects.Debit, operation)

}

func Test_should_be_create_an_invalid_operation(t *testing.T) {
	operation := valueobjects.NewOperation(2)

	require.False(t, operation.IsValid())
	require.False(t, operation.IsCredit())
	require.False(t, operation.IsDebit())

}
