package productEntityTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "otb-order/domain/entity/product"
	"otb-order/test/testDouble/stub/entity/product"
	"testing"
)

type ProductCostShould struct {
	suite.Suite
}

func TestProductCostShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductCostShould))
}

func (s *ProductCostShould) TestCostFrom_ValueIsNegative_ReturnErr() {
	expectedErr := productEntity.ErrCostIsNegative
	negativeCost := productEntityStub.GetCost().Int() * -1

	cost, err := productEntity.CostFrom(negativeCost)

	require.Equal(s.T(), int64(-1), cost.Int())
	errors.Is(expectedErr, err)
}

func (s *ProductCostShould) TestCostFrom_ValueIsValid_ReturnCost() {
	costInt := productEntityStub.GetCost().Int()

	cost, err := productEntity.CostFrom(costInt)

	require.NoError(s.T(), err)
	require.Equal(s.T(), costInt, cost.Int())
}
