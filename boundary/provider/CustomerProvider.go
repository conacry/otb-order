package providerInterface

import (
	"context"
	customerEntity "online-shop-order/domain/entity/customer"
)

type CustomerProvider interface {
	FindByID(ctx context.Context, customerID *customerEntity.CustomerID) (*customerEntity.Customer, error)
}
