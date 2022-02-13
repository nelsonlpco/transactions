package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nelsonlpco/transactions/api/adapter/rest/model"
	"github.com/nelsonlpco/transactions/api/adapter/rest/responses"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/shared/commonerrors"
	"github.com/sirupsen/logrus"
)

type AccountController struct {
	accountService *services.AccountService
}

func NewAccountController(accountService *services.AccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

func (a *AccountController) CreatAccount(w http.ResponseWriter, r *http.Request) {
	var createAccountModel model.CreateAccountModel
	accountId, err := uuid.NewRandom()
	if err != nil {
		logrus.WithField("AccountController", "CreateAccount").Error(err)
		responses.InternalServerError(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&createAccountModel)
	if err != nil {
		logrus.WithField("AccountController", "CreateAccount").Error(err)
		responses.BadRequestResponse(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	accountEntity := entity.NewAccount(accountId, createAccountModel.DocumentNumber)

	err = a.accountService.CreateAccount(r.Context(), accountEntity)
	if err != nil {
		var errorSql *commonerrors.ErrorSql
		var errorInvalidEntity *commonerrors.ErrorInvalidEntity
		if errors.As(err, &errorSql) || errors.As(err, &errorInvalidEntity) {
			responses.BadRequestResponse(w, err.Error())
		} else {
			responses.InternalServerError(w, fmt.Sprintf(`"%v"`, err.Error()))
		}
		return
	}

	shouldReturn := errorManager(err, w, "CreateAccount")
	if shouldReturn {
		return
	}

	responses.SuccessOnCreate(w, fmt.Sprintf(`{"AccountId": "%v"}`, accountId.String()))
	w.WriteHeader(http.StatusCreated)
}

func (a *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId, err := uuid.Parse(params["accountId"])
	if err != nil {
		responses.BadRequestResponse(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	account, err := a.accountService.GetAccountById(r.Context(), accountId)
	shouldReturn := errorManager(err, w, "GetAccount")
	if shouldReturn {
		return
	}

	accountModel := new(model.AccountModel).FromEntity(account)

	payload, err := json.Marshal(accountModel)
	if err != nil {
		responses.InternalServerError(w, err.Error())
		return
	}

	responses.Success(w, string(payload))
}

func errorManager(err error, w http.ResponseWriter, resource string) bool {
	if err != nil {
		var errInvalidEntity *commonerrors.ErrorInvalidEntity
		var errResourceNotFound *commonerrors.ErrorNoContent

		logrus.WithField("AccountController", resource).Error(err)

		if errors.As(err, &errResourceNotFound) {
			responses.NoContent(w)
		} else if errors.As(err, &errInvalidEntity) {
			responses.BadRequestResponse(w, err.Error())
		} else {
			responses.InternalServerError(w, err.Error())
		}

		return true
	}
	return false
}
