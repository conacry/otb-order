package customerEntityTest

import (
	"github.com/conacry/go-platform/pkg/errors"
	"github.com/conacry/go-platform/pkg/generator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerEntity "online-shop-order/domain/entity/customer"
	"testing"
)

type CustomerIDShould struct {
	suite.Suite
}

func TestCustomerIDShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CustomerIDShould))
}

func (s *CustomerIDShould) TestCustomerIDFrom_ValueIsEmpty_ReturnErr() {
	customerID, err := customerEntity.CustomerIDFrom("")
	require.Nil(s.T(), customerID)
	require.ErrorIs(s.T(), err, customerEntity.ErrCustomerIDIsEmpty)
}

func (s *CustomerIDShould) TestCustomerIDFrom_ValueIsNotUUID_ReturnErr() {
	customerID, err := customerEntity.CustomerIDFrom(generator.RandomDefaultStr())
	require.Nil(s.T(), customerID)
	require.True(s.T(), errors.EqualByCode(err, customerEntity.IllegalCustomerIDValueErrCode))
}

func (s *CustomerIDShould) TestCustomerIDFrom_ValueIsValid_ReturnCustomerID() {
	customerIDStr := generator.GenerateUUID()

	customerID, err := customerEntity.CustomerIDFrom(customerIDStr)
	require.NotNil(s.T(), customerID)
	require.Equal(s.T(), customerIDStr, customerID.String())
	require.NoError(s.T(), err)
}
