package orderEntity

import (
	entityCustomer "online-shop-order/internal/domain/entity/customer"
	productEntity "online-shop-order/internal/domain/entity/product"
	"time"
)

type Order struct {
	orderID    *OrderID
	products   []productEntity.Product
	customerID *entityCustomer.CustomerID
	createdAt  *time.Time
}

func (o *Order) OrderID() *OrderID {
	return o.orderID
}

func (o *Order) Products() []productEntity.Product {
	return o.products
}

func (o *Order) CustomerID() *entityCustomer.CustomerID {
	return o.customerID
}

func (o *Order) CreatedAt() *time.Time {
	return o.createdAt
}
