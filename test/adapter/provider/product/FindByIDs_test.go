package productProviderTest

import (
	"context"
	errorsStub "github.com/conacry/go-platform/pkg/errors/test/testDouble/stub"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productProvider "otb-order/adapter/provider/product"
	repositoryMock "otb-order/test/testDouble/mock/repository"
	productEntityStub "otb-order/test/testDouble/stub/entity/product"
	"testing"
)

type FindByIDsShould struct {
	suite.Suite

	ctx             context.Context
	loggerMock      *logMock.LoggerMock
	productRepoMock *repositoryMock.ProductRepositoryMock

	productProvider *productProvider.Provider
}

func TestFindByIDShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(FindByIDsShould))
}

func (s *FindByIDsShould) SetupTest() {
	s.ctx = context.Background()
	s.loggerMock = logMock.GetLogger()
	s.productRepoMock = repositoryMock.GetProductRepository()

	provider, err := productProvider.NewProviderBuilder().
		Logger(s.loggerMock).
		ProductRepository(s.productRepoMock).
		Build()
	require.NoError(s.T(), err)

	s.productProvider = provider
}

func (s *FindByIDsShould) TestFindByIDs_ProductIDsIsNil_ReturnErr() {
	products, err := s.productProvider.FindByIDs(s.ctx, nil)
	assert.Nil(s.T(), products)
	require.ErrorIs(s.T(), err, productProvider.ErrProductIDsIsRequired)
}

func (s *FindByIDsShould) TestFindByIDs_ProductRepoReturnErr_ReturnErr() {
	productIDs := productEntityStub.GetProductIDs(5)

	expectedErr := errorsStub.GetError()
	s.productRepoMock.SetReturnsFor("FindByIDs", expectedErr)

	products, err := s.productProvider.FindByIDs(s.ctx, productIDs)
	assert.Nil(s.T(), products)
	require.ErrorIs(s.T(), err, expectedErr)
}

func (s *FindByIDsShould) TestFindByID_NoErrorsAreOccurred_ReturnCustomer() {
	productIDs := productEntityStub.GetProductIDs(7)

	expectedProducts := productEntityStub.GetProducts(7)
	s.productRepoMock.SetReturnsFor("FindByIDs", expectedProducts)

	products, err := s.productProvider.FindByIDs(s.ctx, productIDs)
	assert.NoError(s.T(), err)
	require.NotEmpty(s.T(), products)

	assert.ElementsMatch(s.T(), expectedProducts, products)
}
