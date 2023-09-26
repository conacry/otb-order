package orderUsecase

import (
	"context"
	log "github.com/conacry/go-platform/pkg/logger"
	boundaryModel "otb-order/boundary/model"
	providerInterface "otb-order/boundary/provider"
	repositoryInterface "otb-order/boundary/repository"
	customerEntity "otb-order/domain/entity/customer"
	orderEntity "otb-order/domain/entity/order"
	productEntity "otb-order/domain/entity/product"
)

var (
	emptyResult = boundaryModel.CreateOrderResult{}
)

type CreateOrder struct {
	logger           log.Logger
	customerProvider providerInterface.CustomerProvider
	productProvider  providerInterface.ProductProvider
	orderRepository  repositoryInterface.OrderRepository
}

func (uc *CreateOrder) Execute(ctx context.Context, params boundaryModel.CreateOrderParams) (boundaryModel.CreateOrderResult, error) {
	if params.IsEmpty() {
		return emptyResult, ErrCreateOrderParamsAreRequired
	}

	customerID, err := customerEntity.CustomerIDFrom(params.CustomerID)
	if err != nil {
		uc.logger.LogError(ctx, err)
		return emptyResult, err
	}

	customer, err := uc.customerProvider.FindByID(ctx, customerID)
	if err != nil {
		uc.logger.LogError(ctx, err)
		return emptyResult, err
	}

	productIDs := make([]*productEntity.ProductID, 0, len(params.ProductIDs))
	for _, productIDStr := range params.ProductIDs {
		productID, err := productEntity.ProductIDFrom(productIDStr)
		if err != nil {
			uc.logger.LogError(ctx, err)
			return emptyResult, err
		}
		productIDs = append(productIDs, productID)
	}

	products, err := uc.productProvider.FindByIDs(ctx, productIDs)
	if err != nil {
		uc.logger.LogError(ctx, err)
		return emptyResult, err
	}

	order, err := orderEntity.CreateOrder(customer, products)
	if err != nil {
		uc.logger.LogError(ctx, err)
		return emptyResult, err
	}

	err = uc.orderRepository.Save(ctx, order)
	if err != nil {
		uc.logger.LogError(ctx, err)
		return emptyResult, err
	}

	result := boundaryModel.CreateOrderResult{
		OrderID: order.OrderID().String(),
	}

	return result, nil
}
