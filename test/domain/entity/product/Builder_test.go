package productEntityTest

import (
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/domain/entity/product"
	productEntityStub "online-shop-order/test/testDouble/stub/entity/product"
	"testing"
)

type ProductBuilderShould struct {
	suite.Suite
}

func TestProductBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductBuilderShould))
}

func (s *ProductBuilderShould) TestNewBuilder_ReturnBuilder() {
	builder := productEntity.NewBuilder()
	require.NotNil(s.T(), builder)
}

func (s *ProductBuilderShould) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		productEntity.ErrProductIDIsRequired,
		productEntity.ErrTitleIsRequired,
		productEntity.ErrDescriptionIsRequired,
		productEntity.ErrCostIsRequired,
	}

	builder, err := productEntity.NewBuilder().Build()
	require.Nil(s.T(), builder)
	commonTesting.AssertErrors(s.T(), err, expectedErr)
}

func (s *ProductBuilderShould) TestBuild_AllParam_ReturnProduct() {
	productID := productEntityStub.GetProductID()
	title := productEntityStub.GetTitle()
	description := productEntityStub.GetDescription()
	cost := productEntityStub.GetCost()

	product, err := productEntity.NewBuilder().
		ProductID(productID).
		Title(title).
		Description(description).
		Cost(cost).
		Build()

	require.NoError(s.T(), err)
	require.NotNil(s.T(), product)
}
