package orderEntityTest

import (
	"github.com/conacry/go-platform/pkg/generator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	orderEntity "online-shop-order/domain/entity/order"
	"testing"
)

type OrderIDShould struct {
	suite.Suite
}

func TestOrderIDShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(OrderIDShould))
}

func (s *OrderIDShould) TestNewOrderID_ReturnOrderID() {
	orderID := orderEntity.NewOrderID()
	require.NotNil(s.T(), orderID)
}

func (s *OrderIDShould) TestOrderIDFrom_ValueIsEmpty_ReturnErr() {
	orderID, err := orderEntity.OrderIDFrom("")
	require.Nil(s.T(), orderID)
	require.ErrorIs(s.T(), err, orderEntity.ErrOrderIDIsEmpty)
}

func (s *OrderIDShould) TestOrderIDFrom_ValueIsNotUUID_ReturnErr() {
	orderID, err := orderEntity.OrderIDFrom(generator.RandomDefaultStr())
	require.Nil(s.T(), orderID)
	require.Error(s.T(), err)
}

func (s *OrderIDShould) TestOrderIDFrom_ValueIsValid_ReturnOrderID() {
	orderIDStr := generator.GenerateUUID()

	orderID, err := orderEntity.OrderIDFrom(orderIDStr)
	require.NoError(s.T(), err)
	require.Equal(s.T(), orderIDStr, orderID.String())
}
