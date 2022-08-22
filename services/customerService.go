package services

import "banking.com/abelh/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetById(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetById(id string) (*domain.Customer, error) {
	return s.repo.GetById(id)
}

func NewCustomerService(repository domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
