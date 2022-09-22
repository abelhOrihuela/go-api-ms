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
	FindAll(status string) ([]Customer, *errs.AppError)
	GetById(id string) (*Customer, *errs.AppError)
}
