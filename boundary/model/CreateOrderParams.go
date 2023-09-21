package boundaryModel

type CreateOrderParams struct {
	CustomerID string
	ProductIDs []string
}

func (p CreateOrderParams) IsEmpty() bool {
	return len(p.CustomerID) == 0 && len(p.ProductIDs) == 0
}
