package customerRepoTest

import (
	"github.com/conacry/go-platform/pkg/generator"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
	mongoMock "github.com/conacry/go-platform/pkg/mongo/test/testDouble/mock"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerRepo "otb-order/adapter/repository/customer"
	"testing"
)

type RepositoryBuilderShould struct {
	suite.Suite
}

func TestRepositoryBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RepositoryBuilderShould))
}

func (s *RepositoryBuilderShould) TestNewRepositoryBuilder_ReturnRepositoryBuilder() {
	builder := customerRepo.NewRepositoryBuilder()
	assert.NotNil(s.T(), builder)
}

func (s *RepositoryBuilderShould) TestBuild_NoParamsWasGiven_ReturnErr() {
	expectedErrs := []error{
		customerRepo.ErrLoggerIsRequired,
		customerRepo.ErrMongoRepoIsRequired,
		customerRepo.ErrCollectionIsRequired,
	}

	repository, err := customerRepo.NewRepositoryBuilder().Build()
	assert.Nil(s.T(), repository)
	commonTesting.AssertErrors(s.T(), err, expectedErrs)
}

func (s *RepositoryBuilderShould) TestBuild_AllParamsWasGiven_ReturnRepository() {
	loggerMock := logMock.GetLogger()
	mongoRepoMock := mongoMock.GetRepository()
	collection, err := mongoModel.CollectionFrom(generator.RandomDefaultStr())
	require.NoError(s.T(), err)

	repository, err := customerRepo.NewRepositoryBuilder().
		Logger(loggerMock).
		MongoRepo(mongoRepoMock).
		Collection(collection).
		Build()
	assert.NoError(s.T(), err)
	require.NotNil(s.T(), repository)

	commonTesting.AssertFieldValue(s.T(), repository, "mongoRepo", mongoRepoMock)
	commonTesting.AssertFieldValue(s.T(), repository, "collection", collection)
	commonTesting.AssertFieldNotNil(s.T(), repository, "errorProcessor")
}
