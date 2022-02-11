package inframodel_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_accountModel_from_account_entity(t *testing.T) {
	accountEntity := entity.NewAccount(uuid.New(), "10094138052")
	accountModel := new(inframodel.AccountModel).FromEntity(accountEntity)

	require.NotNil(t, accountModel)
	require.Equal(t, accountEntity.GetId().String(), accountModel.Id)
	require.Equal(t, accountEntity.GetDocumentNumber(), accountModel.DocumentNumber)
}

func Test_should_be_create_an_accountEntity_from_an_account_model(t *testing.T) {
	accountModel := &inframodel.AccountModel{
		Id:             uuid.NewString(),
		DocumentNumber: "10094138052",
	}
	accountEntity := accountModel.ToEntity()

	require.NotNil(t, accountEntity)
	require.Equal(t, accountModel.Id, accountEntity.GetId().String())
	require.Equal(t, accountModel.DocumentNumber, accountEntity.GetDocumentNumber())
}
