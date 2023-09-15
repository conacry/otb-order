package productEntityStub

import (
	"github.com/conacry/go-platform/pkg/generator"
	productEntity "online-shop-order/domain/entity/product"
)

func GetProductID() *productEntity.ProductID {
	uuidStr := generator.GenerateUUID()
	id, err := productEntity.ProductIDFrom(uuidStr)
	if err != nil {
		panic(err)
	}

	return id
}
