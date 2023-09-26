package repositoryInterface

import (
	"context"
	orderEntity "otb-order/domain/entity/order"
)

type OrderRepository interface {
	Save(ctx context.Context, order *orderEntity.Order) error
}
