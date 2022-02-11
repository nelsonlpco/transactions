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
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
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
		responses.InternalServerError(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}
	transactionId := uuid.New()
	accountId, err := uuid.Parse(createTransactionModel.AccountId)
	if err != nil {
		responses.BadRequestResponse(w, fmt.Sprintf(`"accountId - %v"`, err.Error()))
		return
	}
	operationTypeId, err := uuid.Parse(createTransactionModel.OperationTypeId)
	if err != nil {
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
		var errInvalidEntity *domainerrors.ErrorInvalidEntity

		if errors.As(err, &errInvalidEntity) {
			responses.BadRequestResponse(w, err.Error())
		} else {
			responses.InternalServerError(w, err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}
