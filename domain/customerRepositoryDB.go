package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"banking.com/abelh/errs"
	"banking.com/abelh/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "SELECT * FROM customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "SELECT * FROM customers WHERE status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while scanning table customers" + err.Error())
		return nil, errs.NewUnexpectedDatabaseError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) GetById(id string) (*Customer, *errs.AppError) {
	findAllSql := "SELECT * FROM customers WHERE customer_id = ?"
	var c Customer

	err := d.client.Get(&c, findAllSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning row of table customers" + err.Error())
			return nil, errs.NewUnexpectedDatabaseError("unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDB() CustomerRepositoryDB {

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	passwd := os.Getenv("DB_PASSWD")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, host, port, name)
	// app:An0thrS3crt@tcp(localhost:3306)/banking
	client, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{
		client,
	}
}

// ...
