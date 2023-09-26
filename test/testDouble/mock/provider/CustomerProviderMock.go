package providerMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	customerEntity "otb-order/domain/entity/customer"
)

type CustomerProviderMock struct {
	*mocking.BaseMock
}

func GetCustomerProvider() *CustomerProviderMock {
	return &CustomerProviderMock{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Strict()),
	}
}

func (m *CustomerProviderMock) FindByID(ctx context.Context, customerID *customerEntity.CustomerID) (*customerEntity.Customer, error) {
	res, err := m.ProcessMethod(ctx, customerID)
	if err != nil {
		return nil, err
	}

	if res != nil {
		customer, ok := res.(*customerEntity.Customer)
		if !ok {
			return nil, mocking.ErrCannotCastResult
		}

		return customer, nil
	}

	panic(mocking.ErrImplementMe)
}
