package entity_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_operation_type_enitity(t *testing.T) {
	id := uuid.New()
	operation := valueobjects.NewOperation(valueobjects.Debit)
	operationType := entity.NewOperationType(id, "SAQUE", operation)

	require.NotNil(t, operationType)
	require.Equal(t, id, operationType.GetId())
	require.Equal(t, "SAQUE", operationType.GetDescription())
	require.Equal(t, valueobjects.NewOperation(valueobjects.Debit), operationType.GetOperation())
}

func Test_should_be_create_a_valid_operation_type(t *testing.T) {
	id := uuid.New()
	operation := valueobjects.NewOperation(valueobjects.Credit)
	operationType := entity.NewOperationType(id, "PAGAMENTO", operation)

	err := operationType.Validate()

	require.Nil(t, err)
}

func Test_should_be_create_an_invalid_operation_type_when_operation_is_invalid(t *testing.T) {
	id := uuid.New()
	operation := valueobjects.NewOperation(2)
	operationType := entity.NewOperationType(id, "PAGAMENTO", operation)
	errorMessages := []string{valueobjects.ErrorInvalidOperation.Error()}
	expectedError := commonerrors.NewErrorInvalidEntity("OperationType", errorMessages)

	err := operationType.Validate()

	var errorInvalidEntity *commonerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())

}

func Test_should_be_create_an_invalid_operation_type_when_description_is_empty(t *testing.T) {
	id := uuid.New()
	operation := valueobjects.NewOperation(valueobjects.Credit)
	operationType := entity.NewOperationType(id, "", operation)

	errorMessages := []string{entity.ErrorOperationTypeDescriptionRequired.Error()}
	expectedError := commonerrors.NewErrorInvalidEntity("OperationType", errorMessages)

	err := operationType.Validate()

	var errorInvalidEntity *commonerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}

func Test_should_be_create_an_invalid_operation_type(t *testing.T) {
	id := uuid.New()
	operation := valueobjects.NewOperation(2)

	var errorMessage []string

	errorMessage = append(errorMessage, operation.Validate().Error())
	errorMessage = append(errorMessage, entity.ErrorOperationTypeDescriptionRequired.Error())

	expectedError := commonerrors.NewErrorInvalidEntity("OperationType", errorMessage)

	operationType := entity.NewOperationType(id, "", operation)

	err := operationType.Validate()

	var errorInvalidEntity *commonerrors.ErrorInvalidEntity

	require.True(t, errors.As(err, &errorInvalidEntity))
	require.Equal(t, expectedError.Error(), err.Error())
}
