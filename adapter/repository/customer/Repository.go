package customerRepo

import (
	"github.com/conacry/go-platform/pkg/mongo"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
)

type Repository struct {
	mongoRepo      mongo.Repository
	collection     mongoModel.Collection
	errorProcessor *errorProcessor
}
