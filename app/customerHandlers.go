package app

import (
	"encoding/json"
	"net/http"

	"banking.com/abelh/services"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service services.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	// if only one expected
	filterStatus := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(filterStatus)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customers, err := ch.service.GetById(id)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)

	}
}

func writeResponse(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}
