package customerEntityTest

import (
	"errors"
	"github.com/conacry/go-platform/pkg/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerEntity "online-shop-order/domain/entity/customer"
	"testing"
)

type CustomerProfileShould struct {
	suite.Suite
}

func TestCustomerProfileShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CustomerProfileShould))
}

func (s *CustomerProfileShould) TestProfileFrom_FirstNameIsEmpty_ReturnErr() {
	expectedErr := customerEntity.ErrFirstNameProfileIsEmpty
	lastName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom("", lastName)
	require.Equal(s.T(), customerEntity.Profile{}, profile)
	errors.Is(expectedErr, err)
}

func (s *CustomerProfileShould) TestProfileFrom_LastNameIsEmpty_ReturnErr() {
	expectedErr := customerEntity.ErrLastNameProfileIsEmpty
	firstName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom(firstName, "")
	require.Equal(s.T(), customerEntity.Profile{}, profile)
	errors.Is(expectedErr, err)
}

func (s *CustomerProfileShould) TestProfileFrom_AllParamIsValid_ReturnProfile() {
	firstName := generator.RandomDefaultStr()
	lastName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom(firstName, lastName)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), firstName, profile.FirstName())
	assert.Equal(s.T(), lastName, profile.LastName())
}
