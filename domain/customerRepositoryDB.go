package domain

import (
	"database/sql"
	"time"

	"banking.com/abelh/errs"
	"banking.com/abelh/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "SELECT * FROM customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "SELECT * FROM customers WHERE status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while scanning table customers" + err.Error())
		return nil, errs.NewUnexpectedDatabaseError("unexpected database error")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

		if err != nil {
			logger.Error("Error while scanning row of table customers" + err.Error())
			return nil, errs.NewUnexpectedDatabaseError("unexpected database error")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) GetById(id string) (*Customer, *errs.AppError) {
	findAllSql := "SELECT * FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(findAllSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

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

	client, err := sql.Open("mysql", "app:An0thrS3crt@tcp(localhost:3306)/banking")
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
