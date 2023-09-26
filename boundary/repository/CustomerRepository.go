package repositoryInterface

import (
	"context"
	customerEntity "otb-order/domain/entity/customer"
)

type CustomerRepository interface {
	FindByID(ctx context.Context, customerID *customerEntity.CustomerID) (*customerEntity.Customer, error)
}
