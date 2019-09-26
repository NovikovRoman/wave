package wave

import (
	"github.com/stretchr/testify/require"
	"image"
	"testing"
)

func TestPath_success(t *testing.T) {
	field := NewField([][]float32{
		{0, 0, -1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, -1, 0, 0, 0, 0, 0, 0, 0},
		{-1, 0, 0, 0, -1, 0, 0, 0, 0, 0},
		{0, -1, -1, -1, 0, -1, 0, 0, 0, 0},
		{0, -1, 0, 0, 0, 0, -1, 0, 0, 0},
		{0, -1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	p, ok := Search(field, image.ZP, image.Point{X: 0, Y: 3})
	require.True(t, ok)

	successPath := [][]int{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 3},
		{7, 4},
		{6, 5},
		{5, 5},
		{4, 5},
		{3, 5},
		{2, 5},
		{1, 6},
		{0, 5},
		{0, 4},
		{0, 3},
	}
	for key, point := range *p {
		require.Equal(t, successPath[key][0], point.X)
		require.Equal(t, successPath[key][1], point.Y)
	}
}

func TestPath_notfound(t *testing.T) {
	field := NewField([][]float32{
		{0, 0, -1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, -1, 0, 0, 0, 0, 0, 0, 0},
		{-1, -1, 0, 0, -1, 0, 0, 0, 0, 0},
		{0, -1, -1, -1, 0, -1, 0, 0, 0, 0},
		{0, -1, 0, 0, 0, 0, -1, 0, 0, 0},
		{0, -1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	p, ok := Search(field, image.ZP, image.Point{X: 0, Y: 3})
	require.False(t, ok)
	require.Nil(t, p)
}
