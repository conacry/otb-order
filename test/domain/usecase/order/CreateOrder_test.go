package orderUsecaseTest

import (
	"context"
	"github.com/conacry/go-platform/pkg/errors"
	errorsStub "github.com/conacry/go-platform/pkg/errors/test/testDouble/stub"
	"github.com/conacry/go-platform/pkg/generator"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	mocking "github.com/conacry/go-platform/pkg/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	boundaryModel "online-shop-order/boundary/model"
	usecaseInterface "online-shop-order/boundary/usecase"
	customerEntity "online-shop-order/domain/entity/customer"
	orderEntity "online-shop-order/domain/entity/order"
	productEntity "online-shop-order/domain/entity/product"
	orderUsecase "online-shop-order/domain/usecase/order"
	providerMock "online-shop-order/test/testDouble/mock/provider"
	repositoryMock "online-shop-order/test/testDouble/mock/repository"
	customerEntityStub "online-shop-order/test/testDouble/stub/entity/customer"
	productEntityStub "online-shop-order/test/testDouble/stub/entity/product"
	"testing"
)

type CreateOrderShould struct {
	suite.Suite

	ctx              context.Context
	logger           *logMock.LoggerMock
	customerProvider *providerMock.CustomerProviderMock
	productProvider  *providerMock.ProductProviderMock
	orderRepository  *repositoryMock.OrderRepositoryMock

	createOrderUC usecaseInterface.CreateOrder
}

func TestCreateOrderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CreateOrderShould))
}

func (s *CreateOrderShould) SetupTest() {
	s.logger = logMock.GetLogger()
	s.customerProvider = providerMock.GetCustomerProvider()
	s.productProvider = providerMock.GetProductProvider()
	s.orderRepository = repositoryMock.GetOrderRepository()

	usecase, err := orderUsecase.NewCreateOrderBuilder().
		Logger(s.logger).
		CustomerProvider(s.customerProvider).
		ProductProvider(s.productProvider).
		OrderRepository(s.orderRepository).
		Build()
	require.NoError(s.T(), err)

	s.createOrderUC = usecase
	s.ctx = context.Background()
}

func (s *CreateOrderShould) TestExecute_ParamsIsEmpty_ReturnErr() {
	emptyParams := boundaryModel.CreateOrderParams{}
	require.True(s.T(), emptyParams.IsEmpty())

	result, err := s.createOrderUC.Execute(s.ctx, emptyParams)
	assert.True(s.T(), result.IsEmpty())
	require.ErrorIs(s.T(), err, orderUsecase.ErrCreateOrderParamsAreRequired)
}

func (s *CreateOrderShould) TestExecute_IllegalCustomerID_ReturnErr() {
	params := boundaryModel.CreateOrderParams{
		CustomerID: generator.RandomDefaultStr(),
	}
	require.False(s.T(), params.IsEmpty())

	result, err := s.createOrderUC.Execute(s.ctx, params)
	assert.True(s.T(), result.IsEmpty())
	require.True(s.T(), errors.EqualByCode(err, customerEntity.IllegalCustomerIDValueErrCode))
	assert.True(s.T(), s.logger.IsMethodCalled("LogError"))
}

func (s *CreateOrderShould) TestExecute_CustomerProviderReturnErr_ReturnErr() {
	params := boundaryModel.CreateOrderParams{
		CustomerID: generator.GenerateUUID(),
	}
	require.False(s.T(), params.IsEmpty())

	expectedErr := errorsStub.GetError()
	s.customerProvider.SetReturnsFor("FindByID", expectedErr)

	result, err := s.createOrderUC.Execute(s.ctx, params)
	assert.True(s.T(), result.IsEmpty())
	require.ErrorIs(s.T(), err, expectedErr)
	assert.True(s.T(), s.logger.IsMethodCalled("LogError"))
}

func (s *CreateOrderShould) TestExecute_WrongProductID_ReturnErr() {
	params := boundaryModel.CreateOrderParams{
		CustomerID: generator.GenerateUUID(),
		ProductIDs: []string{
			generator.RandomDefaultStr(),
		},
	}
	require.False(s.T(), params.IsEmpty())

	customer := customerEntityStub.GetCustomer()
	s.customerProvider.SetReturnsFor("FindByID", customer)

	result, err := s.createOrderUC.Execute(s.ctx, params)
	assert.True(s.T(), result.IsEmpty())
	require.True(s.T(), errors.EqualByCode(err, productEntity.IllegalProductIDValueErrCode))
	assert.True(s.T(), s.logger.IsMethodCalled("LogError"))
}

func (s *CreateOrderShould) TestExecute_ProductProviderReturnErr_ReturnErr() {
	params := boundaryModel.CreateOrderParams{
		CustomerID: generator.GenerateUUID(),
		ProductIDs: []string{
			generator.GenerateUUID(),
			generator.GenerateUUID(),
		},
	}
	require.False(s.T(), params.IsEmpty())

	customer := customerEntityStub.GetCustomer()
	s.customerProvider.SetReturnsFor("FindByID", customer)

	expectedErr := errorsStub.GetError()
	s.productProvider.SetReturnsFor("FindByIDs", expectedErr)

	result, err := s.createOrderUC.Execute(s.ctx, params)
	assert.True(s.T(), result.IsEmpty())
	require.ErrorIs(s.T(), err, expectedErr)
	assert.True(s.T(), s.logger.IsMethodCalled("LogError"))
}

func (s *CreateOrderShould) TestExecute_OrderRepositoryReturnErr_ReturnErr() {
	params := boundaryModel.CreateOrderParams{
		CustomerID: generator.GenerateUUID(),
		ProductIDs: []string{
			generator.GenerateUUID(),
			generator.GenerateUUID(),
		},
	}
	require.False(s.T(), params.IsEmpty())

	customer := customerEntityStub.GetCustomer()
	s.customerProvider.SetReturnsFor("FindByID", customer)

	products := productEntityStub.GetProducts(len(params.ProductIDs))
	s.productProvider.SetReturnsFor("FindByIDs", products)

	expectedErr := errorsStub.GetError()
	s.orderRepository.SetReturnsFor("Save", expectedErr)

	result, err := s.createOrderUC.Execute(s.ctx, params)
	assert.True(s.T(), result.IsEmpty())
	require.ErrorIs(s.T(), err, expectedErr)
	assert.True(s.T(), s.logger.IsMethodCalled("LogError"))
}

func (s *CreateOrderShould) TestExecute_NoErrorsAreOccurred_ReturnResult() {
	params := boundaryModel.CreateOrderParams{
		CustomerID: generator.GenerateUUID(),
		ProductIDs: []string{
			generator.GenerateUUID(),
			generator.GenerateUUID(),
		},
	}
	require.False(s.T(), params.IsEmpty())

	customer := customerEntityStub.GetCustomer()
	s.customerProvider.SetReturnsFor("FindByID", customer)

	products := productEntityStub.GetProducts(len(params.ProductIDs))
	s.productProvider.SetReturnsFor("FindByIDs", products)

	s.orderRepository.SetReturnsFor("Save", mocking.NoError)

	result, err := s.createOrderUC.Execute(s.ctx, params)
	assert.False(s.T(), result.IsEmpty())
	require.NoError(s.T(), err)

	args := s.orderRepository.GetCalledArgs("Save")
	require.Len(s.T(), args, 2)

	savedOrder, ok := args[1].(*orderEntity.Order)
	require.True(s.T(), ok)
	assert.Equal(s.T(), savedOrder.OrderID().String(), result.OrderID)

	assert.True(s.T(), s.logger.IsNeverCalled())
}
