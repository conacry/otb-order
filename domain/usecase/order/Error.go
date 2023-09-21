package orderUsecase

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrLoggerIsRequired           = errors.NewError("SYS", "Logger is required")
	ErrCustomerProviderIsRequired = errors.NewError("SYS", "CustomerProvider is required")
	ErrProductProviderIsRequired  = errors.NewError("SYS", "ProductProvider is required")
	ErrOrderRepositoryIsRequired  = errors.NewError("SYS", "OrderRepository is required")
)

var (
	ErrCreateOrderParamsAreRequired = errors.NewError("97ec58df-001", "Params to create an order are required")
)
