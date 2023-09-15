package productEntity

import (
	"github.com/conacry/go-platform/pkg/generator"
)

type ProductID struct {
	value string
}

func ProductIDFrom(strID string) (*ProductID, error) {
	if strID == "" {
		return nil, ErrProductIDIsEmpty
	}

	uuidID, err := generator.UUIDFrom(strID)
	if err != nil {
		return nil, err
	}

	return &ProductID{value: uuidID}, nil
}

func (id *ProductID) String() string {
	return id.value
}
