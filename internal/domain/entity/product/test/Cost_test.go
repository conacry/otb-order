package productTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/internal/domain/entity/product"
	"online-shop-order/internal/domain/entity/product/test/testDouble"
	"testing"
)

type ProductCost struct {
	suite.Suite
}

func TestProductCost(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductCost))
}

func (p *ProductCost) TestCostFrom_ValueIsNegative_ReturnErr() {
	expectedErr := productEntity.ErrCostIsNegative
	negativeCost := testDouble.Cost().Int() * -1

	cost, err := productEntity.CostFrom(negativeCost)

	require.Equal(p.T(), int64(-1), cost.Int())
	errors.Is(expectedErr, err)
}

func (p *ProductCost) TestCostFrom_ValueIsValid_ReturnErr() {
	costInt := testDouble.Cost().Int()

	cost, err := productEntity.CostFrom(costInt)

	require.NoError(p.T(), err)
	require.Equal(p.T(), costInt, cost.Int())
}