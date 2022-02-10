package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nelsonlpco/transactions/api/adapter/rest/model"
	"github.com/nelsonlpco/transactions/application/services"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
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
	var accountModel model.AccountModel

	err := json.NewDecoder(r.Body).Decode(&accountModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountEntity := accountModel.ToEntity()
	accountErrors := accountEntity.Validate()
	if accountErrors != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(accountEntity)

	err = a.accountService.CreateAccount(r.Context(), accountEntity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (a *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId, err := strconv.Atoi(params["accountId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(accountId)

	account, err := a.accountService.GetAccountById(r.Context(), valueobjects.NewId(accountId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accountModel := new(model.AccountModel).FromEntity(account)
	log.Println(accountModel)

	payload, err := json.Marshal(accountModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
