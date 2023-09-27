package customerRepo

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
)

var (
	ErrLoggerIsRequired     = errors.NewError("SYS", "Logger is required")
	ErrMongoRepoIsRequired  = errors.NewError("SYS", "Mongo repository is required")
	ErrCollectionIsRequired = errors.NewError("SYS", "Collection is required")
)

type errorProcessor struct {
	logger log.Logger
}
