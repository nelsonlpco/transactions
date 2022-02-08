package entity_test

import (
	"testing"
	"time"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_transaction(t *testing.T) {
	validDocument := "10094138052"
	account := entity.NewAccount(1, validDocument)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)

	transaction := entity.NewTransaction(1, 100.50, account, operationType, time.Now())

	require.NotNil(t, transaction)
}

func Test_should_be_create_a_valid_transaction(t *testing.T) {
	validDocument := "10094138052"
	account := entity.NewAccount(1, validDocument)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)

	transaction := entity.NewTransaction(1, 100.50, account, operationType, time.Now())

	err := transaction.Validate()

	require.Nil(t, err)
}

func Test_should_be_create_an_invalid_transaction_when_id_is_invalid(t *testing.T) {
	validDocument := "10094138052"
	account := entity.NewAccount(1, validDocument)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)
	expectedError := domainerrors.NewErrorInvalidId("transaction")

	transaction := entity.NewTransaction(0, 100.50, account, operationType, time.Now())

	err := transaction.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, expectedError, err[0])
}

func Test_should_be_create_an_invalid_transaction_when_amount_is_equal_to_zero(t *testing.T) {
	validDocument := "10094138052"
	account := entity.NewAccount(1, validDocument)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)
	expectedError := domainerrors.NewErrorInvalidAmount("transaction")

	transaction := entity.NewTransaction(1, 0, account, operationType, time.Now())

	err := transaction.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, expectedError, err[0])
}

func Test_should_be_create_an_invalid_transaction_when_account_is_invalid(t *testing.T) {
	invalidDocument := "00000138052"
	account := entity.NewAccount(1, invalidDocument)
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)
	expectedError := domainerrors.NewErrorInvalidDocument("account", invalidDocument)

	transaction := entity.NewTransaction(1, 49, account, operationType, time.Now())

	err := transaction.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, expectedError, err[0])
}

func Test_should_be_create_an_invalid_transaction_when_operationType_is_invalid(t *testing.T) {
	validDocument := "10094138052"
	account := entity.NewAccount(1, validDocument)
	operationType := entity.NewOperationType(1, "", valueobjects.Credit)
	expectedError := domainerrors.NewErrorInvalidDescription("operationType")

	transaction := entity.NewTransaction(1, 49, account, operationType, time.Now())

	err := transaction.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, expectedError, err[0])
}

func Test_should_be_create_an_invalid_transaction(t *testing.T) {
	invalidAccount := entity.NewAccount(1, "00000138052")
	invalidOperationType := entity.NewOperationType(1, "", 2)

	transaction := entity.NewTransaction(0, 0, invalidAccount, invalidOperationType, time.Now())

	err := transaction.Validate()

	require.Equal(t, 5, len(err))
}
