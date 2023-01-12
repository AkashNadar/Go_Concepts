package domain

type Customer struct {
	Id      string
	Name    string
	City    string
	Pincode string
	Dob     string
	Status  string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
