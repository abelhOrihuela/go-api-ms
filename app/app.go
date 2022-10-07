package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"banking.com/abelh/domain"
	"banking.com/abelh/logger"
	"banking.com/abelh/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Start() {

	godotenv.Load(".env")

	// define routes and handlers
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	address := fmt.Sprintf("%s:%s", host, port)

	fmt.Print(address)

	router := mux.NewRouter()
	//mux := http.NewServeMux()

	// wiring
	ch := CustomerHandlers{
		// services.NewCustomerService(domain.DefaultCustomerRepositoryStub()),
		services.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getById)

	logger.Info(fmt.Sprintf("Server is running at %s.", address))

	// starting server
	log.Fatal(http.ListenAndServe(address, router))
}
