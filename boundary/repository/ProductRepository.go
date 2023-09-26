package repositoryInterface

import (
	"context"
	productEntity "otb-order/domain/entity/product"
)

type ProductRepository interface {
	FindByIDs(ctx context.Context, productIDs []*productEntity.ProductID) ([]*productEntity.Product, error)
}
