package entity_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_account_entity(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)

	require.NotNil(t, account)
	require.Equal(t, id, account.GetId())
	require.Equal(t, validDocument, account.GetDocumentNumber())
}

func Test_should_be_create_a_valid_account_entity(t *testing.T) {
	account := entity.NewAccount(uuid.New(), "10094138052")

	err := account.Validate()

	require.Nil(t, err)
}

func Test_should_be_create_an_invalid_account_entity_when_document_is_invalid(t *testing.T) {
	invalidDoc := "11100021344"
	account := entity.NewAccount(uuid.New(), invalidDoc)
	invalidDocument := domainerrors.NewErrorInvalidDocument(invalidDoc)
	errorMessages := []string{invalidDocument.Error()}
	expectedError := domainerrors.NewErrorInvalidEntity("Account", errorMessages)

	err := account.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}
