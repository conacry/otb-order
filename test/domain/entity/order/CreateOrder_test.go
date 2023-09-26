package orderEntityTest

import (
	"github.com/conacry/go-platform/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	orderEntity "otb-order/domain/entity/order"
	customerEntityStub "otb-order/test/testDouble/stub/entity/customer"
	productEntityStub "otb-order/test/testDouble/stub/entity/product"
	"testing"
)

type CreateOrderShould struct {
	suite.Suite
}

func TestCreateOrderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CreateOrderShould))
}

func (s *CreateOrderShould) TestCreateOrder_ProductsIsNil_ReturnOrder() {
	customer := customerEntityStub.GetCustomer()

	order, err := orderEntity.CreateOrder(customer, nil)
	require.Nil(s.T(), order)
	require.True(s.T(), errors.ContainByCode(err, orderEntity.ErrProductsIsRequired.Code()))
}

func (s *CreateOrderShould) TestCreateOrder_ProductsIsEmpty_ReturnOrder() {
	emptyProducts := productEntityStub.GetEmptyProducts()
	customer := customerEntityStub.GetCustomer()

	order, err := orderEntity.CreateOrder(customer, emptyProducts)
	require.Nil(s.T(), order)
	require.True(s.T(), errors.ContainByCode(err, orderEntity.ErrProductsIsRequired.Code()))
}

func (s *CreateOrderShould) TestCreateOrder_CustomerIsNil_ReturnOrder() {
	products := productEntityStub.GetProducts(7)

	order, err := orderEntity.CreateOrder(nil, products)
	require.Nil(s.T(), order)
	require.True(s.T(), errors.ContainByCode(err, orderEntity.ErrCustomerIsRequired.Code()))
}

func (s *CreateOrderShould) TestCreateOrder_AllParamIsValid_ReturnOrder() {
	products := productEntityStub.GetProducts(7)
	customer := customerEntityStub.GetCustomer()

	order, err := orderEntity.CreateOrder(customer, products)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), order)

	assert.NotNil(s.T(), order.OrderID())
	assert.NotNil(s.T(), order.CreatedAt())
	assert.EqualValues(s.T(), customer, order.Customer())
	assert.ElementsMatch(s.T(), products, order.Products())
}
