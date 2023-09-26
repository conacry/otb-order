package productProvider

import (
	"context"
	log "github.com/conacry/go-platform/pkg/logger"
	repositoryInterface "otb-order/boundary/repository"
	productEntity "otb-order/domain/entity/product"
)

type Provider struct {
	logger            log.Logger
	productRepository repositoryInterface.ProductRepository
}

func (p *Provider) FindByIDs(ctx context.Context, productIDs []*productEntity.ProductID) ([]*productEntity.Product, error) {
	if productIDs == nil {
		err := ErrProductIDsIsRequired
		p.logger.LogError(ctx, err)
		return nil, err
	}

	return p.productRepository.FindByIDs(ctx, productIDs)
}
