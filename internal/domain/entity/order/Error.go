package orderEntity

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrOrderIDIsEmpty       = errors.NewError("3ee9eb5b-001", "Value to create OrderID is required")
	ErrOrderIDIsRequired    = errors.NewError("3ee9eb5b-002", "OrderID is required")
	ErrProductsIsRequired   = errors.NewError("3ee9eb5b-003", "Products is required")
	ErrCustomerIDIsRequired = errors.NewError("3ee9eb5b-004", "CustomerID is required")
	ErrCreatedAtIsRequired  = errors.NewError("3ee9eb5b-005", "CreatedAt is required")
)
