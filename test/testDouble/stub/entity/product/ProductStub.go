package entityStub

import productEntity "online-shop-order/internal/domain/entity/product"

func Product() *productEntity.Product {
	productID := productEntity.NewProductID()
	title := Title()
	description := Description()
	cost := Cost()

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

func Products(count int) []productEntity.Product {
	products := make([]productEntity.Product, count)
	for i := 0; i < count; i++ {
		product := Product()
		products = append(products, *product)
	}

	return products
}
