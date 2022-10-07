package services

import (
	"banking.com/abelh/domain"
	"banking.com/abelh/dto"
	"banking.com/abelh/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	var response []dto.CustomerResponse

	c, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
