package providerMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	productEntity "online-shop-order/domain/entity/product"
)

type ProductProviderMock struct {
	*mocking.BaseMock
}

func GetProductProvider() *ProductProviderMock {
	return &ProductProviderMock{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Strict()),
	}
}

func (m *ProductProviderMock) FindByIDs(ctx context.Context, productIDs []*productEntity.ProductID) ([]*productEntity.Product, error) {
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
