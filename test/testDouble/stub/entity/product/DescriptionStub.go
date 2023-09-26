package productEntityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "otb-order/domain/entity/product"
)

func GetDescription() *productEntity.Description {
	randomStr := generator.RandomDefaultStr()

	description, err := productEntity.DescriptionFrom(randomStr)
	if err != nil {
		return nil
	}

	return &description
}
