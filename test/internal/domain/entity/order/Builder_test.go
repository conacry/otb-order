package orderEntityTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	entityCustomer "online-shop-order/internal/domain/entity/customer"
	orderEntity "online-shop-order/internal/domain/entity/order"
	"online-shop-order/test/testDouble/stub/entity/product"
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
	products := entityStub.Products(5)
	customerID := entityCustomer.NewCustomerID()
	createdAt := time.Now()

	order, err := orderEntity.NewBuilder().
		OrderID(orderID).
		Products(products).
		CustomerID(customerID).
		CreatedAt(&createdAt).
		Build()

	require.NoError(p.T(), err)
	require.NotNil(p.T(), order)
}
