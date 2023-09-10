package orderTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	entityCustomer "online-shop-order/internal/domain/entity/customer"
	orderEntity "online-shop-order/internal/domain/entity/order"
	"online-shop-order/internal/domain/entity/product/test/testDouble"
	"testing"
	"time"
)

type OrderBuilder struct {
	suite.Suite
}

func TestOrderBuilder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(OrderBuilder))
}

func (o *OrderBuilder) TestNewBuilder_ReturnBuilder() {
	builder := orderEntity.NewBuilder()
	require.NotNil(o.T(), builder)
}

func (o *OrderBuilder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		orderEntity.ErrOrderIDIsRequired,
		orderEntity.ErrProductsIsRequired,
		orderEntity.ErrCustomerIDIsRequired,
		orderEntity.ErrCreatedAtIsRequired,
	}

	builder, err := orderEntity.NewBuilder().Build()
	require.Nil(o.T(), builder)
	commonTesting.AssertErrors(o.T(), err, expectedErr)
}

func (p *OrderBuilder) TestBuild_AllParam_ReturnOrder() {
	orderID := orderEntity.NewOrderID()
	products := testDouble.Products(5)
	customerID := entityCustomer.NewCustomerID()
	createdAt := time.Now()

	builder := orderEntity.NewBuilder()
	builder.OrderID(orderID)
	builder.Products(products)
	builder.CustomerID(customerID)
	builder.CreatedAt(&createdAt)

	order, err := builder.Build()

	require.NoError(p.T(), err)
	require.NotNil(p.T(), order)
}
