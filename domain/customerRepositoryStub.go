package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id: 101010, Name: "Here", ZipCode: "", DateOfBirth: "", Status: "",
		},
	}
	return CustomerRepositoryStub{customers}
}
