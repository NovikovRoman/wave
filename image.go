package wave

import (
	"github.com/NovikovRoman/imglib"
	"image"
	"image/color"
)

func DrawPath(im image.Image, path Path, c color.Color) {
	imSet := im.(imglib.ImageSet)
	for _, p := range path {
		if p.In(im.Bounds()) {
			imSet.Set(p.X, p.Y, c)
		}
	}
}
