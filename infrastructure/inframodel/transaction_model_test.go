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

	transactionModel := new(inframodel.TransactionModel).FromEntity(transactionEntity)

	require.NotNil(t, transactionModel)
	require.Equal(t, transactionEntity.GetId().String(), transactionModel.Id)
	require.Equal(t, float64(transactionEntity.GetAmount()), transactionModel.Amount)
	require.Equal(t, transactionEntity.GetEventDate(), transactionModel.EventDate)
	require.Equal(t, transactionEntity.GetAccount(), transactionModel.Account.ToEntity())
	require.Equal(t, transactionEntity.GetOperationType(), transactionModel.OperationType.ToEntity())
}

func Test_should_be_create_a_transactionTypeEntity_from_a_transactionModel(t *testing.T) {
	accountModel := &inframodel.AccountModel{
		Id:             uuid.NewString(),
		DocumentNumber: "10094138052",
	}

	operationTypeModel := &inframodel.OperationTypeModel{
		Id:          uuid.NewString(),
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}

	transactionModel := &inframodel.TransactionModel{
		Id:            uuid.NewString(),
		Amount:        10.23,
		EventDate:     time.Now(),
		Account:       accountModel,
		OperationType: operationTypeModel,
	}

	transactionEntity := transactionModel.ToEntity()

	require.NotNil(t, transactionEntity)
	require.Equal(t, transactionModel.Id, transactionEntity.GetId().String())
	require.Equal(t, transactionModel.Amount, float64(transactionEntity.GetAmount()))
	require.Equal(t, transactionModel.EventDate, transactionEntity.GetEventDate())
	require.Equal(t, transactionModel.Account.ToEntity(), transactionEntity.GetAccount())
	require.Equal(t, transactionModel.OperationType.ToEntity(), transactionEntity.GetOperationType())
}
