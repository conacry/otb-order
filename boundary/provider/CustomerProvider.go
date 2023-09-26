package providerInterface

import (
	"context"
	customerEntity "otb-order/domain/entity/customer"
)

type CustomerProvider interface {
	FindByID(ctx context.Context, customerID *customerEntity.CustomerID) (*customerEntity.Customer, error)
}
