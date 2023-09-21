package customerEntity

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	IllegalCustomerIDValueErrCode errors.ErrorCode = "d7761015-101"
)

var (
	ErrCustomerIDIsEmpty       = errors.NewError("d7761015-001", "Value to create CustomerID is required")
	ErrFirstNameProfileIsEmpty = errors.NewError("d7761015-002", "Value to create FirstName in Profile is required")
	ErrLastNameProfileIsEmpty  = errors.NewError("d7761015-003", "Value to create LastName in Profile is required")
	ErrCustomerIDIsRequired    = errors.NewError("d7761015-004", "CustomerID is required")
	ErrProfileIsRequired       = errors.NewError("d7761015-005", "Profile is required")
)

func ErrIllegalCustomerIDValue(cause error) error {
	msg := fmt.Sprintf("Value to create CustomerID is illegal. Cause: %q", cause)
	return errors.NewError(IllegalCustomerIDValueErrCode, msg)
}
