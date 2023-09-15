package customerEntityTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerEntity "online-shop-order/domain/entity/customer"
	entityStub "online-shop-order/test/testDouble/stub/entity/customer"
	"testing"
)

type CustomerBuilder struct {
	suite.Suite
}

func TestCustomerBuilder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CustomerBuilder))
}

func (c *CustomerBuilder) TestNewBuilder_ReturnBuilder() {
	builder := customerEntity.NewBuilder()
	require.NotNil(c.T(), builder)
}

func (c *CustomerBuilder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		customerEntity.ErrCustomerIDIsRequired,
		customerEntity.ErrProfileIsRequired,
	}

	builder, err := customerEntity.NewBuilder().Build()

	require.Nil(c.T(), builder)
	commonTesting.AssertErrors(c.T(), err, expectedErr)
}

func (c *CustomerBuilder) TestBuild_AllParam_ReturnCustomer() {
	customerID := customerEntity.NewCustomerID()
	profile := entityStub.Profile()

	customer, err := customerEntity.NewBuilder().
		CustomerID(customerID).
		Profile(profile).
		Build()

	require.NoError(c.T(), err)
	require.NotNil(c.T(), customer)
}
