package domain

import (
	"banking.com/abelh/dto"
	"banking.com/abelh/errs"
)

// Define struct of customers
type Customer struct {
	Id          int64  `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
		ZipCode:     c.ZipCode,
		Status:      statusAsText(c.Status),
		City:        c.City,
	}
}

func statusAsText(status string) string {
	statusText := "active"
	if status == "0" {
		statusText = "inactive"
	}
	return statusText
}

type ICustomerRepository interface {
	/*
		status = 1 -> active
		status = 0 -> inactive
	*/
	FindAll(status string) ([]Customer, *errs.AppError)
	GetById(id string) (*Customer, *errs.AppError)
}
