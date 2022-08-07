package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// define routes and handlers
	host := "localhost:8000"
	router := mux.NewRouter()
	//mux := http.NewServeMux()

	router.HandleFunc("/", helloword).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	fmt.Printf("Server is running at %s.", host)

	// starting server
	log.Fatal(http.ListenAndServe(host, router))
}
