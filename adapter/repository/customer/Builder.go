package customerRepo

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
	"github.com/conacry/go-platform/pkg/mongo"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
)

type RepositoryBuilder struct {
	logger     log.Logger
	mongoRepo  mongo.Repository
	collection mongoModel.Collection

	errors *errors.Errors
}

func NewRepositoryBuilder() *RepositoryBuilder {
	return &RepositoryBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *RepositoryBuilder) Logger(logger log.Logger) *RepositoryBuilder {
	b.logger = logger
	return b
}

func (b *RepositoryBuilder) MongoRepo(mongoRepo mongo.Repository) *RepositoryBuilder {
	b.mongoRepo = mongoRepo
	return b
}

func (b *RepositoryBuilder) Collection(collection mongoModel.Collection) *RepositoryBuilder {
	b.collection = collection
	return b
}

func (b *RepositoryBuilder) Build() (*Repository, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *RepositoryBuilder) checkRequiredFields() {
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
	if b.mongoRepo == nil {
		b.errors.AddError(ErrMongoRepoIsRequired)
	}
	if len(b.collection) == 0 {
		b.errors.AddError(ErrCollectionIsRequired)
	}
}

func (b *RepositoryBuilder) createFromBuilder() *Repository {
	return &Repository{
		mongoRepo:  b.mongoRepo,
		collection: b.collection,
		errorProcessor: &errorProcessor{
			logger: b.logger,
		},
	}
}
