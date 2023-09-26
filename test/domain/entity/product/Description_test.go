package productEntityTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "otb-order/domain/entity/product"
	"otb-order/test/testDouble/stub/entity/product"
	"testing"
)

type ProductDescriptionShould struct {
	suite.Suite
}

func TestProductDescription(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductDescriptionShould))
}

func (s *ProductDescriptionShould) TestDescriptionFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := productEntity.ErrDescriptionIsEmpty

	description, err := productEntity.DescriptionFrom("")

	require.Equal(s.T(), "", description.String())
	errors.Is(expectedErr, err)
}

func (s *ProductDescriptionShould) TestDescriptionFrom_ValueIsValid_ReturnDescription() {
	descriptionStr := productEntityStub.GetDescription().String()

	description, err := productEntity.DescriptionFrom(descriptionStr)

	require.NotNil(s.T(), description)
	require.Equal(s.T(), descriptionStr, description.String())
	require.NoError(s.T(), err)
}
