package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"fullname" xml:"Name"`
	Account string `json:"account" xml:"Account"`
	Level   string `json:"level" xml:"Level"`
}

func helloword(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "¡Hello World!")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	varsRequest := mux.Vars(r)

	customers := []Customer{
		{Name: "Abel", Account: varsRequest["customer_id"], Level: "A"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)

	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}

func createCustomer(rw http.ResponseWriter, r *http.Request) {

}
