package app

import (
	"fmt"
	"log"
	"net/http"

	"banking.com/abelh/domain"
	"banking.com/abelh/logger"
	"banking.com/abelh/services"
	"github.com/gorilla/mux"
)

func Start() {
	// define routes and handlers
	host := "localhost:8000"
	router := mux.NewRouter()
	//mux := http.NewServeMux()

	// wiring
	ch := CustomerHandlers{
		// services.NewCustomerService(domain.DefaultCustomerRepositoryStub()),
		services.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getById)

	logger.Info(fmt.Sprintf("Server is running at %s.", host))

	// starting server
	log.Fatal(http.ListenAndServe(host, router))
}
