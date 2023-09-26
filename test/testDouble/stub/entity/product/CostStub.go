package productEntityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "otb-order/domain/entity/product"
)

func GetCost() *productEntity.Cost {
	randomNumber := generator.RandomNumber(0, 100_000)

	cost, err := productEntity.CostFrom(int64(randomNumber))
	if err != nil {
		return nil
	}

	return &cost
}
