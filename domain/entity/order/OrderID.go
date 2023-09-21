package orderEntity

import "github.com/conacry/go-platform/pkg/generator"

type OrderID struct {
	value string
}

func NewOrderID() *OrderID {
	return &OrderID{value: generator.GenerateUUID()}
}

func OrderIDFrom(strID string) (*OrderID, error) {
	if strID == "" {
		return nil, ErrOrderIDIsEmpty
	}

	uuidID, err := generator.UUIDFrom(strID)
	if err != nil {
		return nil, ErrIllegalOrderIDValue(err)
	}

	return &OrderID{value: uuidID}, nil
}

func (id *OrderID) String() string {
	return id.value
}
