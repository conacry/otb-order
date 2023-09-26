package repositoryMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	orderEntity "otb-order/domain/entity/order"
)

type OrderRepositoryMock struct {
	*mocking.BaseMock
}

func GetOrderRepository() *OrderRepositoryMock {
	return &OrderRepositoryMock{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Strict()),
	}
}

func (m *OrderRepositoryMock) Save(ctx context.Context, order *orderEntity.Order) error {
	_, err := m.ProcessMethod(ctx, order)
	return err
}
