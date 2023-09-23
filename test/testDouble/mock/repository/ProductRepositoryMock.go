package repositoryMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	productEntity "online-shop-order/domain/entity/product"
)

func GetProductRepository() *ProductRepositoryMock {
	return &ProductRepositoryMock{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Strict()),
	}
}

type ProductRepositoryMock struct {
	*mocking.BaseMock
}

func (m *ProductRepositoryMock) FindByIDs(ctx context.Context, productIDs []*productEntity.ProductID) ([]*productEntity.Product, error) {
	res, err := m.ProcessMethod(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	if res != nil {
		products, ok := res.([]*productEntity.Product)
		if !ok {
			return nil, mocking.ErrCannotCastResult
		}

		return products, nil
	}

	panic(mocking.ErrImplementMe)
}
