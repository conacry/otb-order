package orderUsecase

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
	providerInterface "otb-order/boundary/provider"
	repositoryInterface "otb-order/boundary/repository"
)

type CreateOrderBuilder struct {
	logger           log.Logger
	customerProvider providerInterface.CustomerProvider
	productProvider  providerInterface.ProductProvider
	orderRepository  repositoryInterface.OrderRepository

	errors *errors.Errors
}

func NewCreateOrderBuilder() *CreateOrderBuilder {
	return &CreateOrderBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *CreateOrderBuilder) Logger(logger log.Logger) *CreateOrderBuilder {
	b.logger = logger
	return b
}

func (b *CreateOrderBuilder) CustomerProvider(customerProvider providerInterface.CustomerProvider) *CreateOrderBuilder {
	b.customerProvider = customerProvider
	return b
}

func (b *CreateOrderBuilder) ProductProvider(productProvider providerInterface.ProductProvider) *CreateOrderBuilder {
	b.productProvider = productProvider
	return b
}

func (b *CreateOrderBuilder) OrderRepository(orderRepository repositoryInterface.OrderRepository) *CreateOrderBuilder {
	b.orderRepository = orderRepository
	return b
}

func (b *CreateOrderBuilder) Build() (*CreateOrder, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *CreateOrderBuilder) checkRequiredFields() {
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
	if b.customerProvider == nil {
		b.errors.AddError(ErrCustomerProviderIsRequired)
	}
	if b.productProvider == nil {
		b.errors.AddError(ErrProductProviderIsRequired)
	}
	if b.orderRepository == nil {
		b.errors.AddError(ErrOrderRepositoryIsRequired)
	}
}

func (b *CreateOrderBuilder) createFromBuilder() *CreateOrder {
	return &CreateOrder{
		logger:           b.logger,
		customerProvider: b.customerProvider,
		productProvider:  b.productProvider,
		orderRepository:  b.orderRepository,
	}
}
