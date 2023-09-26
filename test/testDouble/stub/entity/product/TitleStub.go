package productEntityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "otb-order/domain/entity/product"
)

func GetTitle() *productEntity.Title {
	randomStr := generator.RandomDefaultStr()

	title, err := productEntity.TitleFrom(randomStr)
	if err != nil {
		return nil
	}

	return &title
}
