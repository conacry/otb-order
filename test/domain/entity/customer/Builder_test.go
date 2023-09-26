package customerEntityTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerEntity "otb-order/domain/entity/customer"
	customerEntityStub "otb-order/test/testDouble/stub/entity/customer"
	"testing"
)

type CustomerBuilderShould struct {
	suite.Suite
}

func TestCustomerBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CustomerBuilderShould))
}

func (s *CustomerBuilderShould) TestNewBuilder_ReturnBuilder() {
	builder := customerEntity.NewBuilder()
	require.NotNil(s.T(), builder)
}

func (s *CustomerBuilderShould) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		customerEntity.ErrCustomerIDIsRequired,
		customerEntity.ErrProfileIsRequired,
	}

	builder, err := customerEntity.NewBuilder().Build()
	require.Nil(s.T(), builder)
	commonTesting.AssertErrors(s.T(), err, expectedErr)
}

func (s *CustomerBuilderShould) TestBuild_AllParam_ReturnCustomer() {
	customerID := customerEntityStub.GetCustomerID()
	profile := customerEntityStub.GetProfile()

	customer, err := customerEntity.NewBuilder().
		CustomerID(customerID).
		Profile(profile).
		Build()
	require.NoError(s.T(), err)
	require.NotNil(s.T(), customer)

	assert.EqualValues(s.T(), customerID, customer.CustomerID())
	assert.EqualValues(s.T(), profile, customer.Profile())
}
