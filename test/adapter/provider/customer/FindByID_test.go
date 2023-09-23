package customerProviderTest

import (
	"context"
	errorsStub "github.com/conacry/go-platform/pkg/errors/test/testDouble/stub"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerProvider "online-shop-order/adapter/provider/customer"
	repositoryMock "online-shop-order/test/testDouble/mock/repository"
	customerEntityStub "online-shop-order/test/testDouble/stub/entity/customer"
	"testing"
)

type FindByIDShould struct {
	suite.Suite

	ctx              context.Context
	loggerMock       *logMock.LoggerMock
	customerRepoMock *repositoryMock.CustomerRepositoryMock

	customerProvider *customerProvider.Provider
}

func TestFindByIDShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(FindByIDShould))
}

func (s *FindByIDShould) SetupTest() {
	s.ctx = context.Background()
	s.loggerMock = logMock.GetLogger()
	s.customerRepoMock = repositoryMock.GetCustomerRepository()

	provider, err := customerProvider.NewProviderBuilder().
		Logger(s.loggerMock).
		CustomerRepository(s.customerRepoMock).
		Build()
	require.NoError(s.T(), err)

	s.customerProvider = provider
}

func (s *FindByIDShould) TestFindByID_CustomerIDIsNil_ReturnErr() {
	customer, err := s.customerProvider.FindByID(s.ctx, nil)
	assert.Nil(s.T(), customer)
	require.ErrorIs(s.T(), err, customerProvider.ErrCustomerIDIsRequired)
}

func (s *FindByIDShould) TestFindByID_CustomerRepoReturnErr_ReturnErr() {
	customerID := customerEntityStub.GetCustomerID()

	expectedErr := errorsStub.GetError()
	s.customerRepoMock.SetReturnsFor("FindByID", expectedErr)

	customer, err := s.customerProvider.FindByID(s.ctx, customerID)
	assert.Nil(s.T(), customer)
	require.ErrorIs(s.T(), err, expectedErr)
}

func (s *FindByIDShould) TestFindByID_NoErrorsAreOccurred_ReturnCustomer() {
	customerID := customerEntityStub.GetCustomerID()

	expectedCustomer := customerEntityStub.GetCustomer()
	s.customerRepoMock.SetReturnsFor("FindByID", expectedCustomer)

	customer, err := s.customerProvider.FindByID(s.ctx, customerID)
	assert.NoError(s.T(), err)
	require.NotNil(s.T(), customer)

	assert.EqualValues(s.T(), expectedCustomer, customer)
}
