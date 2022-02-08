package entity_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_account_entity(t *testing.T) {
	account := entity.NewAccount(1, "10094138052")

	require.NotNil(t, account)
}

func Test_should_be_create_a_valid_account_entity(t *testing.T) {
	account := entity.NewAccount(1, "10094138052")

	err := account.Validate()

	require.Nil(t, err)
}

func Test_should_be_create_an_invalid_account_entity_when_document_is_invalid(t *testing.T) {
	invalidDoc := "11100021344"
	account := entity.NewAccount(1, invalidDoc)
	expectedError := domainerrors.NewErrorInvalidDocument("account", invalidDoc)

	err := account.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, expectedError, err[0])
}

func Test_should_be_create_an_invalid_account_entity_when_id_less_than_or_equal_zero(t *testing.T) {
	invalidIds := []int{-1, 0}

	for _, id := range invalidIds {
		account := entity.NewAccount(id, "10094138052")

		err := account.Validate()

		require.Equal(t, 1, len(err))
		require.Equal(t, domainerrors.NewErrorInvalidId("account"), err[0])
	}
}

func Test_should_be_create_an_invalid_account(t *testing.T) {
	invalidDoc := "11100021344"
	invalidId := 0
	expectedError := domainerrors.NewErrorInvalidDocument("account", invalidDoc)

	account := entity.NewAccount(invalidId, invalidDoc)

	err := account.Validate()

	require.Equal(t, 2, len(err))
	require.Equal(t, expectedError, err[0])
	require.Equal(t, domainerrors.NewErrorInvalidId("account"), err[1])
}
