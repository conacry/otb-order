package customerProvider

import (
	"context"
	log "github.com/conacry/go-platform/pkg/logger"
	repositoryInterface "otb-order/boundary/repository"
	customerEntity "otb-order/domain/entity/customer"
)

type Provider struct {
	logger             log.Logger
	customerRepository repositoryInterface.CustomerRepository
}

func (p *Provider) FindByID(ctx context.Context, customerID *customerEntity.CustomerID) (*customerEntity.Customer, error) {
	if customerID == nil {
		err := ErrCustomerIDIsRequired
		p.logger.LogError(ctx, err)
		return nil, err
	}

	return p.customerRepository.FindByID(ctx, customerID)
}
