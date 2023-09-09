package orderEntity

import (
	"github.com/conacry/go-platform/pkg/errors"
	productEntity "online-shop-order/internal/domain/entity/product"
	"time"
)

type Builder struct {
	orderID    *OrderID
	products   []productEntity.Product
	customerID *CustomerID
	createdAt  *time.Time

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) OrderID(orderID *OrderID) *Builder {
	b.orderID = orderID
	return b
}

func (b *Builder) Products(products []productEntity.Product) *Builder {
	b.products = products
	return b
}

func (b *Builder) CustomerID(customerID *CustomerID) *Builder {
	b.customerID = customerID
	return b
}

func (b *Builder) CreatedAt(createdAt *time.Time) *Builder {
	b.createdAt = createdAt
	return b
}

func (b *Builder) Build() (*Order, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.orderID == nil {
		b.errors.AddError(ErrOrderIDIsRequired)
	}
	if b.products == nil {
		b.errors.AddError(ErrProductsIsRequired)
	}
	if b.customerID == nil {
		b.errors.AddError(ErrCustomerIDIsRequired)
	}
	if b.createdAt == nil {
		b.errors.AddError(ErrCreatedAtIsRequired)
	}
}
func (b *Builder) createFromBuilder() *Order {
	return &Order{
		orderID:    b.orderID,
		products:   b.products,
		customerID: b.customerID,
		createdAt:  b.createdAt,
	}
}
