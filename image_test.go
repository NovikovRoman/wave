package wave

import (
	"github.com/NovikovRoman/imglib"
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"path/filepath"
	"testing"
)

const testdata = "testdata"

func TestDrawPath(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "search.png"))
	require.Nil(t, err)

	imGray := imglib.ToGrayscale(im)
	imglib.ReduceGrayscale(imGray, 2)

	f := NewFieldFromImage(imGray, true, color.Gray{Y: 0})

	p, ok := Search(f,
		image.Point{
			X: 60,
			Y: 0,
		},
		image.Point{
			X: 60,
			Y: im.Bounds().Max.Y - 1,
		})

	require.True(t, ok)

	im  = imglib.ToRGBA(imGray)
	DrawPath(im, *p, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	err = imaging.Save(im, filepath.Join(testdata, "search_result.png"))
	require.Nil(t, err)
}
