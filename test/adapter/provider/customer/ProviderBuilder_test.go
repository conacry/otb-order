package customerProviderTest

import (
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerProvider "online-shop-order/adapter/provider/customer"
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
	builder := customerProvider.NewProviderBuilder()
	assert.NotNil(s.T(), builder)
}

func (s *ProviderBuilderShould) TestBuild_NoParamsIsGiven_ReturnErr() {
	expectedErrs := []error{
		customerProvider.ErrLoggerIsRequired,
		customerProvider.ErrCustomerRepositoryIsRequired,
	}

	provider, err := customerProvider.NewProviderBuilder().Build()
	assert.Nil(s.T(), provider)
	commonTesting.AssertErrors(s.T(), err, expectedErrs)
}

func (s *ProviderBuilderShould) TestBuild_AllParamsIsGiven_ReturnCustomerProvider() {
	loggerMock := logMock.GetLogger()
	customerRepoMock := repositoryMock.GetCustomerRepository()

	provider, err := customerProvider.NewProviderBuilder().
		Logger(loggerMock).
		CustomerRepository(customerRepoMock).
		Build()
	assert.NotNil(s.T(), provider)
	require.NoError(s.T(), err)

	commonTesting.AssertFieldValue(s.T(), provider, "logger", loggerMock)
	commonTesting.AssertFieldValue(s.T(), provider, "customerRepository", customerRepoMock)
}
