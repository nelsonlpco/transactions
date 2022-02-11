package entity_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_transaction(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)
	date := time.Now()
	amount := valueobjects.NewMoney(100.50)

	transaction := entity.NewTransaction(id, amount, account, operationType, date)
	err := transaction.Validate()

	require.Nil(t, err)
	require.NotNil(t, transaction)
	require.Equal(t, operationType, transaction.GetOperationType())
	require.Equal(t, date, transaction.GetEventDate())
	require.Equal(t, account, transaction.GetAccount())
	require.Equal(t, id, transaction.GetId())
	require.Equal(t, amount, transaction.GetAmount())

}

func Test_should_be_create_a_valid_credit_transaction(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)

	amounts := []float64{10.2, -50.0, -20.234}
	expectedAmounts := []valueobjects.Money{10.2, 50.0, 20.234}
	var transaction *entity.Transaction

	for i, amount := range amounts {
		transaction = entity.NewTransaction(uuid.New(), valueobjects.NewMoney(amount), account, operationType, time.Now())
		err := transaction.Validate()

		require.Equal(t, expectedAmounts[i], transaction.GetAmount())
		require.Nil(t, err)
	}
}

func Test_should_be_create_a_valid_debit_transaction(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)
	operationType := entity.NewOperationType(id, "COMPRAS", valueobjects.Debit)

	amounts := []float64{10.2, -50.0, 12}
	expectedAmounts := []valueobjects.Money{-10.2, -50.0, -12}
	var transaction *entity.Transaction

	for i, amount := range amounts {
		transaction = entity.NewTransaction(uuid.New(), valueobjects.NewMoney(amount), account, operationType, time.Now())
		err := transaction.Validate()

		require.Equal(t, expectedAmounts[i], transaction.GetAmount())
		require.Nil(t, err)
	}
}

func Test_should_be_create_an_invalid_transaction_when_amount_is_equal_to_zero(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)

	errorMessages := []string{valueobjects.ErrorInvalidAmmount.Error()}
	expectedError := domainerrors.NewErrorInvalidEntity("Transaction", errorMessages)

	transaction := entity.NewTransaction(id, valueobjects.NewMoney(0), account, operationType, time.Now())

	err := transaction.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}

func Test_should_be_create_an_invalid_transaction_when_account_is_nil(t *testing.T) {
	id := uuid.New()
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)

	errorMessages := []string{entity.ErrorTransactionAccountNotBeNil.Error()}
	expectedError := domainerrors.NewErrorInvalidEntity("Transaction", errorMessages)

	transaction := entity.NewTransaction(id, 49, nil, operationType, time.Now())

	err := transaction.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())

}

func Test_should_be_create_an_invalid_transaction_when_account_is_invalid(t *testing.T) {
	id := uuid.New()
	operationType := entity.NewOperationType(id, "PAGAMENTO", valueobjects.Credit)
	account := entity.NewAccount(id, "000011123")

	errorMesages := []string{account.Validate().Error()}
	expectedError := domainerrors.NewErrorInvalidEntity("Transaction", errorMesages)

	transaction := entity.NewTransaction(id, 49, account, operationType, time.Now())

	err := transaction.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())

}

func Test_should_be_create_an_invalid_transaction_when_operationType_is_nil(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)

	errorMessages := []string{entity.ErrorTransactionOperationTypeNotBeNil.Error()}
	expectedError := domainerrors.NewErrorInvalidEntity("Transaction", errorMessages)

	transaction := entity.NewTransaction(id, 49, account, nil, time.Now())

	err := transaction.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}
func Test_should_be_create_an_invalid_transaction_when_operationType_is_invalid(t *testing.T) {
	validDocument := "10094138052"
	id := uuid.New()
	account := entity.NewAccount(id, validDocument)
	operationType := entity.NewOperationType(id, "", 2)

	errorMessages := []string{operationType.Validate().Error()}
	expectedError := domainerrors.NewErrorInvalidEntity("Transaction", errorMessages)

	transaction := entity.NewTransaction(id, 49, account, operationType, time.Now())

	err := transaction.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}

func Test_should_be_create_an_invalid_transaction(t *testing.T) {
	id := uuid.New()
	invalidAccount := entity.NewAccount(id, "00000138052")
	invalidOperationType := entity.NewOperationType(id, "", 2)

	errorMessages := []string{
		invalidAccount.Validate().Error(),
		invalidOperationType.Validate().Error(),
		valueobjects.ErrorInvalidAmmount.Error(),
	}

	expectedError := domainerrors.NewErrorInvalidEntity("Transaction", errorMessages)

	transaction := entity.NewTransaction(id, 0, invalidAccount, invalidOperationType, time.Now())

	err := transaction.Validate()

	var errorInvalidEntity *domainerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}
