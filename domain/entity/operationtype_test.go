package entity_test

import (
	"testing"

	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_operation_type_enitity(t *testing.T) {
	operationType := entity.NewOperationType(1, "SAQUE", valueobjects.Debit)

	require.NotNil(t, operationType)
}

func Test_should_be_create_a_valid_operation_type(t *testing.T) {
	operationType := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)

	err := operationType.Validate()

	require.Nil(t, err)
}

func Test_should_be_create_an_invalid_operation_type_when_id_is_invalid(t *testing.T) {
	operationType := entity.NewOperationType(0, "PAGAMENTO", valueobjects.Credit)

	err := operationType.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, domainerrors.NewErrorInvalidId("operationType"), err[0])
}

func Test_should_be_create_an_invalid_operation_type_when_operation_is_invalid(t *testing.T) {
	operationType := entity.NewOperationType(1, "PAGAMENTO", 2)

	err := operationType.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, domainerrors.NewErrorInvalidOperation("operationType"), err[0])
}

func Test_should_be_create_an_invalid_operation_type_when_description_is_empty(t *testing.T) {
	operationType := entity.NewOperationType(1, "", valueobjects.Credit)

	err := operationType.Validate()

	require.Equal(t, 1, len(err))
	require.Equal(t, domainerrors.NewErrorInvalidDescription("operationType"), err[0])
}

func Test_should_be_create_an_invalid_operation_type(t *testing.T) {
	operationType := entity.NewOperationType(0, "", 2)

	err := operationType.Validate()

	require.Equal(t, 3, len(err))
	require.Equal(t, domainerrors.NewErrorInvalidId("operationType"), err[0])
	require.Equal(t, domainerrors.NewErrorInvalidOperation("operationType"), err[1])
	require.Equal(t, domainerrors.NewErrorInvalidDescription("operationType"), err[2])
}