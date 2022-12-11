package app

import (
	"encoding/json"
	"net/http"

	"banking.com/abelh/dto"
	"banking.com/abelh/services"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service services.AccountService
}

func (ah *AccountHandlers) NewAccount(rw http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	customerId := requestVars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)

		if appError != nil {
			writeResponse(rw, appError.Code, appError.Message)
		} else {
			writeResponse(rw, http.StatusCreated, account)
		}

	}
}
