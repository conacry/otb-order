package customerProvider

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
	repositoryInterface "otb-order/boundary/repository"
)

type ProviderBuilder struct {
	logger             log.Logger
	customerRepository repositoryInterface.CustomerRepository

	errors *errors.Errors
}

func NewProviderBuilder() *ProviderBuilder {
	return &ProviderBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *ProviderBuilder) Logger(logger log.Logger) *ProviderBuilder {
	b.logger = logger
	return b
}

func (b *ProviderBuilder) CustomerRepository(customerRepository repositoryInterface.CustomerRepository) *ProviderBuilder {
	b.customerRepository = customerRepository
	return b
}

func (b *ProviderBuilder) Build() (*Provider, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *ProviderBuilder) checkRequiredFields() {
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
	if b.customerRepository == nil {
		b.errors.AddError(ErrCustomerRepositoryIsRequired)
	}
}

func (b *ProviderBuilder) createFromBuilder() *Provider {
	return &Provider{
		logger:             b.logger,
		customerRepository: b.customerRepository,
	}
}
