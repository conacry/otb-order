package productTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/internal/domain/entity/product"
	"online-shop-order/internal/domain/entity/product/test/testDouble"
	"testing"
)

type ProductBuilder struct {
	suite.Suite
}

func TestProductBuilder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductBuilder))
}

func (p *ProductBuilder) TestNewBuilder_ReturnBuilder() {
	builder := productEntity.NewBuilder()
	require.NotNil(p.T(), builder)
}

func (p *ProductBuilder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		productEntity.ErrProductIDIsRequired,
		productEntity.ErrTitleIsRequired,
		productEntity.ErrDescriptionIsRequired,
		productEntity.ErrCostIsRequired,
	}

	builder, err := productEntity.NewBuilder().Build()
	require.Nil(p.T(), builder)
	commonTesting.AssertErrors(p.T(), err, expectedErr)
}

func (p *ProductBuilder) TestBuild_AllParam_ReturnProduct() {
	productID := productEntity.NewProductID()
	title := testDouble.Title()
	description := testDouble.Description()
	cost := testDouble.Cost()

	builder := productEntity.NewBuilder()
	builder.ProductID(productID)
	builder.Title(title)
	builder.Description(description)
	builder.Cost(cost)

	product, err := builder.Build()

	require.NoError(p.T(), err)
	require.NotNil(p.T(), product)
}
