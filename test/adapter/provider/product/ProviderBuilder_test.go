package productProviderTest

import (
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productProvider "online-shop-order/adapter/provider/product"
	repositoryMock "online-shop-order/test/testDouble/mock/repository"
	"testing"
)

type ProviderBuilderShould struct {
	suite.Suite
}

func TestProviderBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProviderBuilderShould))
}

func (s *ProviderBuilderShould) TestNewProviderBuilder_ReturnProviderBuilder() {
	builder := productProvider.NewProviderBuilder()
	assert.NotNil(s.T(), builder)
}

func (s *ProviderBuilderShould) TestBuild_NoParamsIsGiven_ReturnErr() {
	expectedErrs := []error{
		productProvider.ErrLoggerIsRequired,
		productProvider.ErrProductRepositoryIsRequired,
	}

	provider, err := productProvider.NewProviderBuilder().Build()
	assert.Nil(s.T(), provider)
	commonTesting.AssertErrors(s.T(), err, expectedErrs)
}

func (s *ProviderBuilderShould) TestBuild_AllParamsIsGiven_ReturnproductProvider() {
	loggerMock := logMock.GetLogger()
	productRepository := repositoryMock.GetProductRepository()

	provider, err := productProvider.NewProviderBuilder().
		Logger(loggerMock).
		ProductRepository(productRepository).
		Build()
	assert.NotNil(s.T(), provider)
	require.NoError(s.T(), err)

	commonTesting.AssertFieldValue(s.T(), provider, "logger", loggerMock)
	commonTesting.AssertFieldValue(s.T(), provider, "productRepository", productRepository)
}
