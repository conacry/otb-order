package customerEntity

type Customer struct {
	customerID *CustomerID
	profile    *Profile
}

func (c Customer) CustomerID() *CustomerID {
	return c.customerID
}

func (c Customer) Profile() *Profile {
	return c.profile
}
