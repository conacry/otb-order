package customerProvider

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrLoggerIsRequired             = errors.NewError("SYS", "Logger is required")
	ErrCustomerRepositoryIsRequired = errors.NewError("SYS", "Customer repository is required")
)

var (
	ErrCustomerIDIsRequired = errors.NewError("acfb604a-001", "CustomerIDIsRequired")
)
