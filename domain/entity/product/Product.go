package productEntity

type Product struct {
	productID   *ProductID
	title       *Title
	description *Description
	cost        *Cost
}

func (p *Product) ProductID() *ProductID {
	return p.productID
}

func (p *Product) Title() *Title {
	return p.title
}

func (p *Product) Description() *Description {
	return p.description
}

func (p *Product) Cost() *Cost {
	return p.cost
}
