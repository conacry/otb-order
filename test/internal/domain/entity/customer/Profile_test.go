package customerEntityTest

import (
	"errors"
	"fmt"
	"github.com/conacry/go-platform/pkg/generator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customerEntity "online-shop-order/internal/domain/entity/customer"
	"testing"
)

type CustomerProfile struct {
	suite.Suite
}

func TestCustomerProfile(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CustomerProfile))
}

func (c *CustomerProfile) TestProfileFrom_FirstNameIsEmpty_ReturnErr() {
	expectedErr := customerEntity.ErrFirstNameProfileIsEmpty
	lastName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom("", lastName)

	require.Equal(c.T(), customerEntity.Profile{}, profile)
	errors.Is(expectedErr, err)
}

func (c *CustomerProfile) TestProfileFrom_LastNameIsEmpty_ReturnErr() {
	expectedErr := customerEntity.ErrLastNameProfileIsEmpty
	firstName := generator.RandomDefaultStr()

	profile, err := customerEntity.ProfileFrom(firstName, "")

	require.Equal(c.T(), customerEntity.Profile{}, profile)
	errors.Is(expectedErr, err)
}

func (c *CustomerProfile) TestProfileFrom_AllParamIsValid_ReturnProfile() {
	firstName := generator.RandomDefaultStr()
	lastName := generator.RandomDefaultStr()
	profileStr := firstName + " " + lastName

	profile, err := customerEntity.ProfileFrom(firstName, lastName)

	fmt.Println(profileStr, profile.String())
	require.Equal(c.T(), profileStr, profile.String())
	require.NoError(c.T(), err)
}
