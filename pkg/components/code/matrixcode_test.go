// nolint: dupl
package code_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/test"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
)

func TestNewMatrix(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewMatrix("code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewMatrix("code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_custom_prop.json")
	})
}

func TestTestNewMatrixCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewMatrixCol(12, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewMatrixCol(12, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_col_custom_prop.json")
	})
}

func TestNewMatrixRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := code.NewMatrixRow(10, "code")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := code.NewMatrixRow(10, "code", fixture.RectProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_row_custom_prop.json")
	})
}

func TestMatrixCode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := code.NewMatrix(codeValue, prop)

		provider := &mocks.Provider{}
		provider.EXPECT().AddMatrixCode(codeValue, &cell, &prop)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddMatrixCode", 1)
	})
}

func TestMatrixCode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		sut := code.NewMatrix("code")

		// Act
		sut.SetConfig(nil)
	})
}
