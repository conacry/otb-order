package productEntity

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	IllegalProductIDValueErrCode errors.ErrorCode = "16f7d28d-101"
)

var (
	ErrProductIDIsEmpty      = errors.NewError("16f7d28d-001", "Value to create ProductID is required")
	ErrTitleIsEmpty          = errors.NewError("16f7d28d-002", "Title cannot be empty")
	ErrTitleTooLong          = errors.NewError("16f7d28d-003", "Title is too long")
	ErrDescriptionIsEmpty    = errors.NewError("16f7d28d-004", "Description cannot be empty")
	ErrDescriptionTooLong    = errors.NewError("16f7d28d-005", "Description is too long")
	ErrCostIsNegative        = errors.NewError("16f7d28d-006", "Cost cannot be negative")
	ErrProductIDIsRequired   = errors.NewError("16f7d28d-007", "ProductID is required")
	ErrTitleIsRequired       = errors.NewError("16f7d28d-008", "Product title is required")
	ErrDescriptionIsRequired = errors.NewError("16f7d28d-009", "Product description is required")
	ErrCostIsRequired        = errors.NewError("16f7d28d-010", "Product cost is required")
)

func ErrIllegalProductIDValue(cause error) error {
	msg := fmt.Sprintf("Value to create OrderID is illegal. Cause: %q", cause)
	return errors.NewError(IllegalProductIDValueErrCode, msg)
}
