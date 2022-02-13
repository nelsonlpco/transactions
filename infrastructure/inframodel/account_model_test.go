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
	accountModel, err := new(inframodel.AccountModel).FromEntity(accountEntity)

	expectedId := uuid.New()
	expectedId.UnmarshalBinary(accountModel.Id)

	require.Nil(t, err)
	require.NotNil(t, accountModel)
	require.Equal(t, accountEntity.GetId(), expectedId)
	require.Equal(t, accountEntity.GetDocumentNumber(), accountModel.DocumentNumber)
}

func Test_should_be_create_an_accountEntity_from_an_account_model(t *testing.T) {
	accountId, _ := uuid.NewRandom()
	binaryId, _ := accountId.MarshalBinary()

	accountModel := &inframodel.AccountModel{
		Id:             binaryId,
		DocumentNumber: "10094138052",
	}
	accountEntity, err := accountModel.ToEntity()

	require.Nil(t, err)
	require.NotNil(t, accountEntity)
	require.Equal(t, accountId, accountEntity.GetId())
	require.Equal(t, accountModel.DocumentNumber, accountEntity.GetDocumentNumber())
}
