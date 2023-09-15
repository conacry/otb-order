package productEntityTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/internal/domain/entity/product"
	"testing"
)

type ProductID struct {
	suite.Suite
}

func TestProductID(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductID))
}

func (p *ProductID) TestNewProductID_ReturnProductID() {
	productID := productEntity.NewProductID()
	require.NotNil(p.T(), productID)
}

func (p *ProductID) TestProductIDFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := productEntity.ErrProductIDIsEmpty
	productID, err := productEntity.ProductIDFrom("")

	require.Nil(p.T(), productID)
	errors.Is(expectedErr, err)
}

func (p *ProductID) TestProductIDFrom_AllParam_ReturnProductID() {
	productIDStr := productEntity.NewProductID().String()

	productID, err := productEntity.ProductIDFrom(productIDStr)

	require.NoError(p.T(), err)
	require.NotNil(p.T(), productID)
}
