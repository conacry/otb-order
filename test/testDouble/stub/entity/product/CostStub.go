package entityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "online-shop-order/internal/domain/entity/product"
)

func Cost() *productEntity.Cost {
	randomNumber := generator.RandomNumber(0, 100_000)

	cost, err := productEntity.CostFrom(int64(randomNumber))
	if err != nil {
		return nil
	}

	return &cost
}
