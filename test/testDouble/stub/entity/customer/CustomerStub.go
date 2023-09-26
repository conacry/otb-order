package customerEntityStub

import customerEntity "otb-order/domain/entity/customer"

func GetCustomer() *customerEntity.Customer {
	customer, err := customerEntity.NewBuilder().
		CustomerID(GetCustomerID()).
		Profile(GetProfile()).
		Build()
	if err != nil {
		panic(err)
	}

	return customer
}
