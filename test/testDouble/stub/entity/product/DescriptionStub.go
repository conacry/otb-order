package entityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "online-shop-order/domain/entity/product"
)

func Description() *productEntity.Description {
	randomStr := generator.RandomDefaultStr()

	description, err := productEntity.DescriptionFrom(randomStr)
	if err != nil {
		return nil
	}

	return &description
}
