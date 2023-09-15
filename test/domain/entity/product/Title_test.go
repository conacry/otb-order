package productEntityTest

import (
	"github.com/conacry/go-platform/pkg/generator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productEntity "online-shop-order/domain/entity/product"
	"testing"
)

type TitleShould struct {
	suite.Suite
}

func TestTitleShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TitleShould))
}

func (s *TitleShould) TestTitleFrom_ValueIsEmpty_ReturnErr() {
	expectedErr := productEntity.ErrTitleIsEmpty

	title, err := productEntity.TitleFrom("")
	require.Empty(s.T(), title)
	require.ErrorIs(s.T(), expectedErr, err)
}

func (s *TitleShould) TestTitleFrom_ValueIsTooLong_ReturnTitle() {
	titleStr := generator.RandomStr(350)

	title, err := productEntity.TitleFrom(titleStr)
	require.Empty(s.T(), title)
	require.ErrorIs(s.T(), err, productEntity.ErrTitleTooLong)
}

func (s *TitleShould) TestTitleFrom_ValueIsValid_ReturnTitle() {
	titleStr := generator.RandomStr(100)
	title, err := productEntity.TitleFrom(titleStr)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), title)
	require.Equal(s.T(), titleStr, title.String())
}
