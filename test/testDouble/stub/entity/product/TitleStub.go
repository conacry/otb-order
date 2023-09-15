package entityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "online-shop-order/domain/entity/product"
)

func Title() *productEntity.Title {
	randomStr := generator.RandomDefaultStr()

	title, err := productEntity.TitleFrom(randomStr)
	if err != nil {
		return nil
	}

	return &title
}
