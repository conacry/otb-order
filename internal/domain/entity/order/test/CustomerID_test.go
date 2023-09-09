package orderTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	orderEntity "online-shop-order/internal/domain/entity/order"
	"testing"
)

type CustomerID struct {
	suite.Suite
}

func TestCustomerID(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CustomerID))
}

func (c *CustomerID) TestNewCustomerID_ReturnCustomerID() {
	customerID := orderEntity.NewCustomerID()
	require.NotNil(c.T(), customerID)
}

func (c *CustomerID) TestCustomerIDFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := orderEntity.ErrCustomerIDIsEmpty

	customerID, err := orderEntity.CustomerIDFrom("")

	require.Nil(c.T(), customerID)
	errors.Is(expectedErr, err)
}

func (c *CustomerID) TestCustomerIDFrom_ValueIsValid_ReturnErr() {
	customerIDStr := orderEntity.NewCustomerID().String()

	customerID, err := orderEntity.CustomerIDFrom(customerIDStr)

	require.NotNil(c.T(), customerID)
	require.Equal(c.T(), customerIDStr, customerID.String())
	require.NoError(c.T(), err)
}
