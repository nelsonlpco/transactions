package inframodel_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_operationTypeModel_from_operationTypeEntity(t *testing.T) {
	operationTypeEntity := entity.NewOperationType(uuid.New(), "PAGAMENTO", valueobjects.Credit)
	operationTypeModel, err := new(inframodel.OperationTypeModel).FromEntity(operationTypeEntity)

	expectedId := uuid.New()
	expectedId.UnmarshalBinary(operationTypeModel.Id)

	require.Nil(t, err)
	require.NotNil(t, operationTypeModel)
	require.Equal(t, operationTypeEntity.GetId(), expectedId)
	require.Equal(t, operationTypeEntity.GetDescription(), operationTypeModel.Description)
	require.Equal(t, byte(operationTypeEntity.GetOperation()), operationTypeModel.Operation)
}

func Test_should_be_create_an_operationTypeEntity_from_an_operationTypeModel(t *testing.T) {
	id := uuid.New()
	binaryId, _ := id.MarshalBinary()
	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          binaryId,
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}
	operationTypeEntity, err := operationTypeModel.ToEntity()

	require.Nil(t, err)
	require.NotNil(t, operationTypeModel)
	require.Equal(t, id, operationTypeEntity.GetId())
	require.Equal(t, operationTypeModel.Description, operationTypeEntity.GetDescription())
	require.Equal(t, operationTypeModel.Operation, byte(operationTypeEntity.GetOperation()))
}
