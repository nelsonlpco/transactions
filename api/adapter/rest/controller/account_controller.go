package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nelsonlpco/transactions/api/adapter/rest/model"
	"github.com/nelsonlpco/transactions/api/adapter/rest/responses"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/domainerrors"
	"github.com/nelsonlpco/transactions/domain/entity"
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
	accountId, _ := uuid.NewRandom()
	id, _ := accountId.Value()
	b, _ := accountId.MarshalBinary()
	log.Println("-------------")
	logrus.Info(accountId)
	logrus.Info(id)
	logrus.Info(b)
	log.Println("-------------")

	err := json.NewDecoder(r.Body).Decode(&createAccountModel)
	if err != nil {
		responses.BadRequestResponse(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	accountEntity := entity.NewAccount(accountId, createAccountModel.DocumentNumber)

	err = a.accountService.CreateAccount(r.Context(), accountEntity)
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
func (a *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId, err := uuid.Parse(params["accountId"])
	if err != nil {
		responses.BadRequestResponse(w, fmt.Sprintf(`"%v"`, err.Error()))
		return
	}

	account, err := a.accountService.GetAccountById(r.Context(), accountId)
	if err != nil {
		var errInvalidEntity *domainerrors.ErrorInvalidEntity

		if errors.As(err, &errInvalidEntity) {
			responses.BadRequestResponse(w, err.Error())
		} else {
			responses.InternalServerError(w, err.Error())
		}

		return
	}

	accountModel := new(model.AccountModel).FromEntity(account)

	payload, err := json.Marshal(accountModel)
	if err != nil {
		responses.InternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
