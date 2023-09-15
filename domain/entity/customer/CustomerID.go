package customerEntity

import (
	"github.com/conacry/go-platform/pkg/generator"
)

type CustomerID struct {
	value string
}

func CustomerIDFrom(strID string) (*CustomerID, error) {
	if strID == "" {
		return nil, ErrCustomerIDIsEmpty
	}

	uuidID, err := generator.UUIDFrom(strID)
	if err != nil {
		return nil, err
	}

	return &CustomerID{value: uuidID}, nil
}

func (id *CustomerID) String() string {
	return id.value
}
