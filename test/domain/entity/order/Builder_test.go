package orderEntityTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	orderEntity "online-shop-order/domain/entity/order"
	customerEntityStub "online-shop-order/test/testDouble/stub/entity/customer"
	"online-shop-order/test/testDouble/stub/entity/product"
	"testing"
	"time"
)

type OrderBuilderShould struct {
	suite.Suite
}

func TestOrderBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(OrderBuilderShould))
}

func (s *OrderBuilderShould) TestNewBuilder_ReturnBuilder() {
	builder := orderEntity.NewBuilder()
	require.NotNil(s.T(), builder)
}

func (s *OrderBuilderShould) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		orderEntity.ErrOrderIDIsRequired,
		orderEntity.ErrProductsIsRequired,
		orderEntity.ErrCustomerIDIsRequired,
		orderEntity.ErrCreatedAtIsRequired,
	}

	builder, err := orderEntity.NewBuilder().Build()
	require.Nil(s.T(), builder)
	commonTesting.AssertErrors(s.T(), err, expectedErr)
}

func (s *OrderBuilderShould) TestBuild_AllParam_ReturnOrder() {
	orderID := orderEntity.NewOrderID()
	products := productEntityStub.GetProducts(5)
	customerID := customerEntityStub.GetCustomerID()
	createdAt := time.Now()

	order, err := orderEntity.NewBuilder().
		OrderID(orderID).
		Products(products).
		CustomerID(customerID).
		CreatedAt(&createdAt).
		Build()
	require.NoError(s.T(), err)
	require.NotNil(s.T(), order)
}
