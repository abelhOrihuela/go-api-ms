package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"banking.com/abelh/services"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"fullname" xml:"Name"`
	Account string `json:"account" xml:"Account"`
	Level   string `json:"level" xml:"Level"`
}

type CustomerHandlers struct {
	service services.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}

func (ch *CustomerHandlers) getById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customers, err := ch.service.GetById(id)

	if err != nil {
		rw.WriteHeader(err.Code)
		fmt.Fprintf(rw, err.Message)
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			rw.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(rw).Encode(customers)
		} else {
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(customers)
		}
	}
}
