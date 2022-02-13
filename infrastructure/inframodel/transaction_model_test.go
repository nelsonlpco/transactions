package inframodel_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_a_transactionModel_from_transactionEntity(t *testing.T) {
	accountEntity := entity.NewAccount(uuid.New(), "10094138052")
	operationTypeEntity := entity.NewOperationType(uuid.New(), "PAGAMENTO", valueobjects.Credit)
	transactionEntity := entity.NewTransaction(
		uuid.New(),
		valueobjects.NewMoney(10.23),
		accountEntity,
		operationTypeEntity,
		time.Now(),
	)

	transactionModel, err := new(inframodel.TransactionModel).FromEntity(transactionEntity)

	account, _ := transactionModel.Account.ToEntity()
	operationType, _ := transactionModel.OperationType.ToEntity()
	expectedId, _ := transactionEntity.GetId().MarshalBinary()

	require.Nil(t, err)
	require.NotNil(t, transactionModel)
	require.Equal(t, expectedId, transactionModel.Id)
	require.Equal(t, float64(transactionEntity.GetAmount()), transactionModel.Amount)
	require.Equal(t, transactionEntity.GetEventDate(), transactionModel.EventDate)
	require.Equal(t, transactionEntity.GetAccount(), account)
	require.Equal(t, transactionEntity.GetOperationType(), operationType)
}

func Test_should_be_create_a_transactionTypeEntity_from_a_transactionModel(t *testing.T) {
	accountId := uuid.New()
	binaryAccoutId, _ := accountId.MarshalBinary()
	accountModel := &inframodel.AccountModel{
		Id:             binaryAccoutId,
		DocumentNumber: "10094138052",
	}

	operationId := uuid.New()
	binaryOperationId, _ := operationId.MarshalBinary()
	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          binaryOperationId,
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}

	id := uuid.New()
	binaryId, _ := id.MarshalBinary()
	transactionModel := &inframodel.TransactionModel{
		Id:            binaryId,
		Amount:        10.23,
		EventDate:     time.Now(),
		Account:       accountModel,
		OperationType: operationTypeModel,
	}

	transactionEntity, err := transactionModel.ToEntity()

	account, _ := transactionModel.Account.ToEntity()
	operationType, _ := transactionModel.OperationType.ToEntity()

	require.Nil(t, err)
	require.NotNil(t, transactionEntity)
	require.Equal(t, id, transactionEntity.GetId())
	require.Equal(t, transactionModel.Amount, float64(transactionEntity.GetAmount()))
	require.Equal(t, transactionModel.EventDate, transactionEntity.GetEventDate())
	require.Equal(t, account, transactionEntity.GetAccount())
	require.Equal(t, operationType, transactionEntity.GetOperationType())
}
