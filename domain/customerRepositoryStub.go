package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func DefaultCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id: 101010, Name: "Here", ZipCode: "", DateOfBirth: "", Status: "",
		},
		{
			Id: 101070, Name: "Here 2m-", ZipCode: "", DateOfBirth: "", Status: "",
		},
	}
	return CustomerRepositoryStub{customers}
}
