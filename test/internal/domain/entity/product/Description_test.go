package productEntityTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/internal/domain/entity/product"
	"online-shop-order/test/testDouble/stub/entity/product"
	"testing"
)

type ProductDescription struct {
	suite.Suite
}

func TestProductDescription(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductDescription))
}

func (p *ProductDescription) TestDescriptionFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := productEntity.ErrDescriptionIsEmpty

	description, err := productEntity.DescriptionFrom("")

	require.Equal(p.T(), "", description.String())
	errors.Is(expectedErr, err)
}

func (p *ProductDescription) TestDescriptionFrom_ValueIsValid_ReturnDescription() {
	descriptionStr := entityStub.Description().String()

	description, err := productEntity.DescriptionFrom(descriptionStr)

	require.NotNil(p.T(), description)
	require.Equal(p.T(), descriptionStr, description.String())
	require.NoError(p.T(), err)
}
