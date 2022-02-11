package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func Test_should_be_create_an_opreation_type_repository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	operationDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
	operationRepository := repository.NewOperationTypeRepositoryImpl(operationDatasource)

	require.NotNil(t, operationRepository)
}

func Test_CreateOperationTypeHandler(t *testing.T) {
	validId := uuid.New()
	creditOperationEntity := entity.NewOperationType(validId, "PAGAMENTO", valueobjects.Credit)
	creditOperationModel := &inframodel.OperationTypeModel{
		Id:          validId.String(),
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}
	rootCtx := context.Background()

	t.Run("should be create a valid operaton type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operationDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
		operationRepository := repository.NewOperationTypeRepositoryImpl(operationDatasource)

		operationDatasource.EXPECT().Create(rootCtx, creditOperationModel).Return(nil)

		err := operationRepository.Create(rootCtx, creditOperationEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an error when operationTypeDatasource fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operationDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
		operationRepository := repository.NewOperationTypeRepositoryImpl(operationDatasource)
		expectedError := operationRepository.MakeError("fail")

		operationDatasource.EXPECT().Create(rootCtx, creditOperationModel).Return(errors.New("fail"))

		err := operationRepository.Create(rootCtx, creditOperationEntity)

		require.Equal(t, expectedError, err)
	})
}

func Test_GetOperationTypeByIdHandler(t *testing.T) {
	validId := uuid.New()
	creditOperationEntity := entity.NewOperationType(validId, "PAGAMENTO", valueobjects.Credit)
	creditOperationModel := &inframodel.OperationTypeModel{
		Id:          validId.String(),
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}
	rootCtx := context.Background()

	t.Run("should be get a valid operaton type by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operationDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
		operationRepository := repository.NewOperationTypeRepositoryImpl(operationDatasource)

		operationDatasource.EXPECT().GetById(rootCtx, validId.String()).Return(creditOperationModel, nil)

		operationType, err := operationRepository.GetById(rootCtx, validId)

		require.Nil(t, err)
		require.Equal(t, creditOperationEntity, operationType)
	})

	t.Run("should be return an error when operationTypeDatasource fail to get an operation type by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operationDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
		operationRepository := repository.NewOperationTypeRepositoryImpl(operationDatasource)
		expectedError := operationRepository.MakeError("fail")

		operationDatasource.EXPECT().GetById(rootCtx, validId.String()).Return(nil, errors.New("fail"))

		operationType, err := operationRepository.GetById(rootCtx, validId)

		require.Nil(t, operationType)
		require.Equal(t, expectedError, err)
	})
}
