package initApp

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrAppContextIsRequired = errors.NewError("SYS", "App context is required")
	ErrAppConfigIsRequired  = errors.NewError("SYS", "App config is required")
	ErrIllegalHttpPort      = errors.NewError("SYS", "Illegal value of http port")
)
