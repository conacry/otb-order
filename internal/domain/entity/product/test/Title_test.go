package productTest

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/internal/domain/entity/product"
	"online-shop-order/internal/domain/entity/product/test/testDouble"
	"testing"
)

type ProductTitle struct {
	suite.Suite
}

func TestProductTitle(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductTitle))
}

func (p *ProductTitle) TestTitleFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := productEntity.ErrTitleIsEmpty

	title, err := productEntity.TitleFrom("")

	require.Equal(p.T(), "", title.String())
	errors.Is(expectedErr, err)
}

func (p *ProductTitle) TestTitleFrom_ValueIsValid_ReturnTitle() {
	titleStr := testDouble.Title().String()

	title, err := productEntity.TitleFrom(titleStr)

	require.NoError(p.T(), err)
	require.NotNil(p.T(), title)
	require.Equal(p.T(), titleStr, title.String())
}
