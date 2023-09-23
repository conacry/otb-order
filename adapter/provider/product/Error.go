package productProvider

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrLoggerIsRequired            = errors.NewError("SYS", "Logger is required")
	ErrProductRepositoryIsRequired = errors.NewError("SYS", "Product repository is required")
)

var (
	ErrProductIDsIsRequired = errors.NewError("e26baa36-001", "ProductIDsIsRequired")
)
