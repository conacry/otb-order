package orderEntity

import (
	"github.com/conacry/go-platform/pkg/errors"
	"github.com/conacry/go-platform/pkg/time"
	entityCustomer "online-shop-order/domain/entity/customer"
	productEntity "online-shop-order/domain/entity/product"
)

type Builder struct {
	orderID   *OrderID
	products  []*productEntity.Product
	customer  *entityCustomer.Customer
	createdAt *time.Time

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

func (b *Builder) Products(products []*productEntity.Product) *Builder {
	b.products = products
	return b
}

func (b *Builder) Customer(customer *entityCustomer.Customer) *Builder {
	b.customer = customer
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
	if len(b.products) == 0 {
		b.errors.AddError(ErrProductsIsRequired)
	}
	if b.customer == nil {
		b.errors.AddError(ErrCustomerIsRequired)
	}
	if b.createdAt == nil {
		b.errors.AddError(ErrCreatedAtIsRequired)
	}
}
func (b *Builder) createFromBuilder() *Order {
	return &Order{
		orderID:   b.orderID,
		products:  b.products,
		customer:  b.customer,
		createdAt: b.createdAt,
	}
}
