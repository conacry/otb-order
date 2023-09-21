package providerInterface

import (
	"context"
	productEntity "online-shop-order/domain/entity/product"
)

type ProductProvider interface {
	FindByIDs(ctx context.Context, productIDs []*productEntity.ProductID) ([]*productEntity.Product, error)
}
