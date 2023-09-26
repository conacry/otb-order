package orderEntity

import (
	"github.com/conacry/go-platform/pkg/time"
	entityCustomer "otb-order/domain/entity/customer"
	productEntity "otb-order/domain/entity/product"
)

type Order struct {
	orderID   *OrderID
	products  []*productEntity.Product
	customer  *entityCustomer.Customer
	createdAt *time.Time
}

func CreateOrder(customer *entityCustomer.Customer, products []*productEntity.Product) (*Order, error) {
	orderID := NewOrderID()
	createdAt := time.Now()

	return NewBuilder().
		OrderID(orderID).
		Customer(customer).
		Products(products).
		CreatedAt(createdAt).
		Build()
}

func (o *Order) OrderID() *OrderID {
	return o.orderID
}

func (o *Order) Products() []*productEntity.Product {
	return o.products
}

func (o *Order) Customer() *entityCustomer.Customer {
	return o.customer
}

func (o *Order) CreatedAt() *time.Time {
	return o.createdAt
}
