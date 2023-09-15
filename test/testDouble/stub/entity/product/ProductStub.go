package productEntityStub

import productEntity "online-shop-order/domain/entity/product"

func GetProduct() *productEntity.Product {
	productID := GetProductID()
	title := GetTitle()
	description := GetDescription()
	cost := GetCost()

	builder := productEntity.NewBuilder()
	builder.ProductID(productID)
	builder.Title(title)
	builder.Description(description)
	builder.Cost(cost)

	product, err := builder.Build()
	if err != nil {
		return nil
	}

	return product
}

func GetProducts(count int) []productEntity.Product {
	products := make([]productEntity.Product, count)
	for i := 0; i < count; i++ {
		product := GetProduct()
		products = append(products, *product)
	}

	return products
}
