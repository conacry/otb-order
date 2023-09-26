package customerEntityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	customerEntity "otb-order/domain/entity/customer"
)

func GetCustomerID() *customerEntity.CustomerID {
	uuidStr := generator.GenerateUUID()
	id, err := customerEntity.CustomerIDFrom(uuidStr)
	if err != nil {
		panic(err)
	}

	return id
}
