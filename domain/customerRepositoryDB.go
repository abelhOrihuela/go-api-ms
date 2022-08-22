package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSql := "SELECT * FROM customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying table customers", err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning row of table customers", err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) GetById(id string) (*Customer, error) {
	findAllSql := "SELECT * FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(findAllSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

	if err != nil {
		log.Println("Error while scanning row of table customers", err.Error())
		return nil, err
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
