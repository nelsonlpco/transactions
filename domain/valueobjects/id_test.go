package valueobjects_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_id(t *testing.T) {
	id := valueobjects.NewId(1)

	require.NotNil(t, id)
}

func Test_should_be_create_an_valid_id(t *testing.T) {
	id := valueobjects.NewId(1)

	require.True(t, id.IsValid())
	require.Equal(t, 1, id.ToInt())

}

func Test_should_be_create_a_invalid_id(t *testing.T) {
	id := valueobjects.NewId(0)

	require.False(t, id.IsValid())
	require.Equal(t, 0, id.ToInt())

}
