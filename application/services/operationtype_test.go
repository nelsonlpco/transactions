package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/usecases"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	mock_datasource "github.com/nelsonlpco/transactions/infrastructure/datasource/mock"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/nelsonlpco/transactions/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func operationTypeServiceCreateInfra(t *testing.T) (*gomock.Controller, *services.OperationTypeService, *mock_datasource.MockOperationTypeDatasource) {
	ctrl := gomock.NewController(t)
	operationTypeDatasource := mock_datasource.NewMockOperationTypeDatasource(ctrl)
	operationTypeRepository := repository.NewOperationTypeRepositoryImpl(operationTypeDatasource)
	createOperatinTypeUseCase := usecases.NewCreateOperationType(operationTypeRepository)
	getOperationTypeByIdUseCase := usecases.NewGetOperationTypeById(operationTypeRepository)
	operationTypeService := services.NewOperationTypeService(getOperationTypeByIdUseCase, createOperatinTypeUseCase)

	return ctrl, operationTypeService, operationTypeDatasource
}

func Test_should_be_create_an_operation_type_service(t *testing.T) {
	ctrl, operationTypeService, _ := operationTypeServiceCreateInfra(t)
	defer ctrl.Finish()

	require.NotNil(t, operationTypeService)
}

func Test_CreateOperationTypeServiceHandler(t *testing.T) {
	operationModel := &inframodel.OperationTypeModel{
		Id:          1,
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}

	operationTypeEntity := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)

	t.Run("should be create an operation type", func(t *testing.T) {
		ctrl, operationTypeService, operationTypeDatasource := operationTypeServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		operationTypeDatasource.EXPECT().Create(ctx, operationModel).Return(nil)

		err := operationTypeService.CreateOperationType(ctx, operationTypeEntity)

		require.Nil(t, err)
	})

	t.Run("should be return an error when data source fail on create the operation type", func(t *testing.T) {
		ctrl, operationTypeService, operationTypeDatasource := operationTypeServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		operationTypeDatasource.EXPECT().Create(ctx, operationModel).Return(errors.New("fail"))

		err := operationTypeService.CreateOperationType(ctx, operationTypeEntity)

		require.NotNil(t, err)
	})

	t.Run("should be return an error when data receive an invalid operation type", func(t *testing.T) {
		ctrl, operationTypeService, _ := operationTypeServiceCreateInfra(t)
		defer ctrl.Finish()

		invalidOperationTypeEntity := entity.NewOperationType(1, "", valueobjects.Credit)
		ctx := context.Background()

		err := operationTypeService.CreateOperationType(ctx, invalidOperationTypeEntity)

		require.NotNil(t, err)
	})
}

func Test_GetOperationTypeByIdServiceHandler(t *testing.T) {
	operationModel := &inframodel.OperationTypeModel{
		Id:          1,
		Description: "PAGAMENTO",
		Operation:   byte(valueobjects.Credit),
	}

	operationTypeEntity := entity.NewOperationType(1, "PAGAMENTO", valueobjects.Credit)

	t.Run("should be get an operation type by id", func(t *testing.T) {
		ctrl, operationTypeService, operationTypeDatasource := operationTypeServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		operationTypeDatasource.EXPECT().GetById(ctx, 1).Return(operationModel, nil)

		operationType, err := operationTypeService.GetOperationTypeById(ctx, 1)

		require.Nil(t, err)
		require.Equal(t, operationTypeEntity, operationType)
	})

	t.Run("should be return an error when data source fail to get an operation type by id", func(t *testing.T) {
		ctrl, operationTypeService, operationTypeDatasource := operationTypeServiceCreateInfra(t)
		defer ctrl.Finish()

		ctx := context.Background()
		operationTypeDatasource.EXPECT().GetById(ctx, 1).Return(nil, errors.New("fail"))

		operationType, err := operationTypeService.GetOperationTypeById(ctx, 1)

		require.Nil(t, operationType)
		require.NotNil(t, err)
	})

	t.Run("should be return an error when data receive an invalid operation type", func(t *testing.T) {
		ctrl, operationTypeService, _ := operationTypeServiceCreateInfra(t)
		defer ctrl.Finish()

		invalidOperationTypeEntity := entity.NewOperationType(1, "", valueobjects.Credit)
		ctx := context.Background()

		err := operationTypeService.CreateOperationType(ctx, invalidOperationTypeEntity)

		require.NotNil(t, err)
	})
}
