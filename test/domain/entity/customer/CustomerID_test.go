package customerEntityTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	entityCustomer "online-shop-order/domain/entity/customer"
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
	customerID := entityCustomer.NewCustomerID()
	require.NotNil(c.T(), customerID)
}

func (c *CustomerID) TestCustomerIDFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := entityCustomer.ErrCustomerIDIsEmpty

	customerID, err := entityCustomer.CustomerIDFrom("")

	require.Nil(c.T(), customerID)
	errors.Is(expectedErr, err)
}

func (c *CustomerID) TestCustomerIDFrom_ValueIsValid_ReturnCustomerID() {
	customerIDStr := entityCustomer.NewCustomerID().String()

	customerID, err := entityCustomer.CustomerIDFrom(customerIDStr)

	require.NotNil(c.T(), customerID)
	require.Equal(c.T(), customerIDStr, customerID.String())
	require.NoError(c.T(), err)
}
