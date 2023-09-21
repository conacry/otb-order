package boundaryModel

type CreateOrderResult struct {
	OrderID string
}

func (r CreateOrderResult) IsEmpty() bool {
	return len(r.OrderID) == 0
}
