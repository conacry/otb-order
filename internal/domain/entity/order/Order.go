package orderEntity

import (
	productEntity "online-shop-order/internal/domain/entity/product"
	"time"
)

type Order struct {
	orderID    *OrderID
	products   []productEntity.Product
	customerID *CustomerID
	createdAt  *time.Time
}

func (o *Order) OrderID() *OrderID {
	return o.orderID
}

func (o *Order) Products() []productEntity.Product {
	return o.products
}

func (o *Order) CustomerID() *CustomerID {
	return o.customerID
}

func (o *Order) CreatedAt() *time.Time {
	return o.createdAt
}
