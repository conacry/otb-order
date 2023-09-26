package productEntityTest

import (
	"errors"
	"github.com/conacry/go-platform/pkg/generator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "otb-order/domain/entity/product"
	"testing"
)

type ProductIDShould struct {
	suite.Suite
}

func TestProductIDShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductIDShould))
}

func (s *ProductIDShould) TestProductIDFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := productEntity.ErrProductIDIsEmpty

	productID, err := productEntity.ProductIDFrom("")
	require.Nil(s.T(), productID)
	errors.Is(expectedErr, err)
}

func (s *ProductIDShould) TestProductIDFrom_ValueIsNotUUID_ReturnErr() {
	productID, err := productEntity.ProductIDFrom(generator.RandomDefaultStr())
	require.Nil(s.T(), productID)
	require.Error(s.T(), err)
}

func (s *ProductIDShould) TestProductIDFrom_AllParam_ReturnProductID() {
	productIDStr := generator.GenerateUUID()

	productID, err := productEntity.ProductIDFrom(productIDStr)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), productID)
}
