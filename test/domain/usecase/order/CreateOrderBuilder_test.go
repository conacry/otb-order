package orderUsecaseTest

import (
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	orderUsecase "otb-order/domain/usecase/order"
	providerMock "otb-order/test/testDouble/mock/provider"
	repositoryMock "otb-order/test/testDouble/mock/repository"
	"testing"
)

type CreateOrderBuilderShould struct {
	suite.Suite
}

func TestCreateOrderBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CreateOrderBuilderShould))
}

func (s *CreateOrderBuilderShould) TestNewCreateOrderBuilder_ReturnCreateOrderBuilder() {
	builder := orderUsecase.NewCreateOrderBuilder()
	assert.NotNil(s.T(), builder)
}

func (s *CreateOrderBuilderShould) TestBuild_NoParams_ReturnErr() {
	expectedErrs := []error{
		orderUsecase.ErrLoggerIsRequired,
		orderUsecase.ErrCustomerProviderIsRequired,
		orderUsecase.ErrProductProviderIsRequired,
		orderUsecase.ErrOrderRepositoryIsRequired,
	}

	useCase, err := orderUsecase.NewCreateOrderBuilder().Build()
	assert.Nil(s.T(), useCase)
	commonTesting.AssertErrors(s.T(), err, expectedErrs)
}

func (s *CreateOrderBuilderShould) TestBuild_AllParams_ReturnCreateOrderUseCase() {
	loggerMock := logMock.GetLogger()
	customerProviderMock := providerMock.GetCustomerProvider()
	productProviderMock := providerMock.GetProductProvider()
	orderRepoMock := repositoryMock.GetOrderRepository()

	useCase, err := orderUsecase.NewCreateOrderBuilder().
		Logger(loggerMock).
		CustomerProvider(customerProviderMock).
		ProductProvider(productProviderMock).
		OrderRepository(orderRepoMock).
		Build()
	assert.NotNil(s.T(), useCase)
	require.NoError(s.T(), err)

	commonTesting.AssertFieldValue(s.T(), useCase, "logger", loggerMock)
	commonTesting.AssertFieldValue(s.T(), useCase, "customerProvider", customerProviderMock)
	commonTesting.AssertFieldValue(s.T(), useCase, "productProvider", productProviderMock)
	commonTesting.AssertFieldValue(s.T(), useCase, "orderRepository", orderRepoMock)
}
