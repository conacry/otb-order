package orderEntity

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	IllegalOrderIDValueErrCode errors.ErrorCode = "3ee9eb5b-101"
)

var (
	ErrOrderIDIsEmpty      = errors.NewError("3ee9eb5b-001", "Value to create OrderID is required")
	ErrOrderIDIsRequired   = errors.NewError("3ee9eb5b-002", "OrderID is required")
	ErrProductsIsRequired  = errors.NewError("3ee9eb5b-003", "Products is required")
	ErrCustomerIsRequired  = errors.NewError("3ee9eb5b-004", "Customer is required")
	ErrCreatedAtIsRequired = errors.NewError("3ee9eb5b-005", "CreatedAt is required")
)

func ErrIllegalOrderIDValue(cause error) error {
	msg := fmt.Sprintf("Value to create OrderID is illegal. Cause: %q", cause)
	return errors.NewError(IllegalOrderIDValueErrCode, msg)
}
