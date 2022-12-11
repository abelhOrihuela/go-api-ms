package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"banking.com/abelh/domain"
	"banking.com/abelh/logger"
	"banking.com/abelh/services"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	if os.Getenv("APP_PORT") == "" ||
		os.Getenv("APP_HOST") == "" {
		log.Fatal("Environment variables not defined... ")
	}
}

func Start() {

	godotenv.Load(".env")
	sanityCheck()

	// define routes and handlers
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	address := fmt.Sprintf("%s:%s", host, port)

	fmt.Print(address)

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// Initialize client of DB
	dbClient := getDbClient()

	// customerRepository
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccounRepositoryDB(dbClient)

	// wiring
	ch := CustomerHandlers{
		//services.NewCustomerService(domain.DefaultCustomerRepositoryStub()),
		services.NewCustomerService(customerRepositoryDB),
	}

	ah := AccountHandlers{
		services.NewAccountService(accountRepositoryDB),
	}

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getById)

	logger.Info(fmt.Sprintf("Server is running at %s.", address))

	// starting server
	log.Fatal(http.ListenAndServe(address, router))
}

func getDbClient() *sqlx.DB {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	passwd := os.Getenv("DB_PASSWD")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, host, port, name)
	client, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
