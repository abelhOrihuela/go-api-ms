package domain

import "banking.com/abelh/errs"

// Define struct of customers
type Customer struct {
	Id          int64
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}
type ICustomerRepository interface {
	FindAll() ([]Customer, error)
	GetById(string) (*Customer, *errs.AppError)
}
