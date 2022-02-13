package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonlpco/transactions/api/adapter/rest/model"
	"github.com/nelsonlpco/transactions/api/adapter/rest/responses"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

func (t *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var createTransactionModel model.CreateTransactionModel
	err := json.NewDecoder(r.Body).Decode(&createTransactionModel)
	if err != nil {
		logrus.WithField("TransactionController", "CreateTransaction").Error(err)
		responses.BadRequestResponse(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	transactionId, err := uuid.NewRandom()
	if err != nil {
		logrus.WithField("TransactionController", "CreateTransaction").Error(err)
		responses.BadRequestResponse(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	accountId, err := uuid.Parse(createTransactionModel.AccountId)
	if err != nil {
		logrus.WithField("TransactionController", "CreateTransaction").Error(err)
		responses.BadRequestResponse(w, fmt.Sprintf(`"accountId - %v"`, err.Error()))
		return
	}

	operationTypeId, err := uuid.Parse(createTransactionModel.OperationTypeId)
	if err != nil {
		logrus.WithField("TransactionController", "CreateTransaction").Error(err)
		responses.BadRequestResponse(w, fmt.Sprintf(`"operatonTypeId - %v"`, err.Error()))
		return
	}

	err = t.transactionService.CreateTransaction(
		r.Context(),
		transactionId,
		accountId,
		operationTypeId,
		valueobjects.NewMoney(createTransactionModel.Amount),
		time.Now(),
	)

	if err != nil {
		var errInvalidEntity *commonerrors.ErrorInvalidEntity
		var errResourceNotFound *commonerrors.ErrorNoContent

		if errors.As(err, &errResourceNotFound) || errors.As(err, &errInvalidEntity) {
			responses.BadRequestResponse(w, err.Error())
		} else {
			responses.InternalServerError(w, err.Error())
		}
		return
	}

	responses.SuccessOnCreate(w, fmt.Sprintf(`{"transactionId": "%v"}`, transactionId.String()))
}
