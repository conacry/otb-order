package orderEntity

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrOrderIDIsEmpty       = errors.NewError("3ee9eb5b-001", "Value to create OrderID is required")
	ErrCustomerIDIsEmpty    = errors.NewError("3ee9eb5b-002", "Value to create CustomerID is required")
	ErrOrderIDIsRequired    = errors.NewError("3ee9eb5b-003", "OrderID is required")
	ErrProductsIsRequired   = errors.NewError("3ee9eb5b-004", "Products is required")
	ErrCustomerIDIsRequired = errors.NewError("3ee9eb5b-005", "CustomerID is required")
	ErrCreatedAtIsRequired  = errors.NewError("3ee9eb5b-006", "CreatedAt is required")
)
