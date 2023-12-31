package usecaseInterface

import (
	"context"
	boundaryModel "otb-order/boundary/model"
)

type CreateOrder interface {
	Execute(ctx context.Context, params boundaryModel.CreateOrderParams) (boundaryModel.CreateOrderResult, error)
}
