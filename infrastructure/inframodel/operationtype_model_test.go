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
	operationTypeModel := new(inframodel.OperationTypeModel).FromEntity(operationTypeEntity)

	require.NotNil(t, operationTypeModel)
	require.Equal(t, operationTypeEntity.GetId().String(), operationTypeModel.Id)
	require.Equal(t, operationTypeEntity.GetDescription(), operationTypeModel.Description)
	require.Equal(t, byte(operationTypeEntity.GetOperation()), operationTypeModel.Operation)
}

func Test_should_be_create_an_operationTypeEntity_from_an_operationTypeModel(t *testing.T) {
	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          uuid.NewString(),
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}
	operationTypeEntity := operationTypeModel.ToEntity()

	require.NotNil(t, operationTypeModel)
	require.Equal(t, operationTypeModel.Id, operationTypeEntity.GetId().String())
	require.Equal(t, operationTypeModel.Description, operationTypeEntity.GetDescription())
	require.Equal(t, operationTypeModel.Operation, byte(operationTypeEntity.GetOperation()))
}
