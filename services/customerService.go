package services

import (
	"banking.com/abelh/domain"
	"banking.com/abelh/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetById(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.GetById(id)
}

func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
