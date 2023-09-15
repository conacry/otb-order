package orderEntityTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	orderEntity "online-shop-order/domain/entity/order"
	"testing"
)

type OrderID struct {
	suite.Suite
}

func TestOrderID(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(OrderID))
}

func (o *OrderID) TestNewOrderID_ReturnOrderID() {
	orderID := orderEntity.NewOrderID()
	require.NotNil(o.T(), orderID)
}

func (o *OrderID) TestOrderIDFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := orderEntity.ErrOrderIDIsEmpty

	orderID, err := orderEntity.OrderIDFrom("")

	require.Nil(o.T(), orderID)
	errors.Is(expectedErr, err)
}

func (o *OrderID) TestOrderIDFrom_ValueIsValid_ReturnOrderID() {
	orderIDStr := orderEntity.NewOrderID().String()

	orderID, err := orderEntity.OrderIDFrom(orderIDStr)

	require.NoError(o.T(), err)
	require.Equal(o.T(), orderIDStr, orderID.String())
}
