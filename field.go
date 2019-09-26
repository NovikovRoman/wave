package wave

import (
	"image"
	"image/color"
)

const Obstacle = -1
const Empty = 0

type IField interface {
	Width() int
	Height() int
	Value(point image.Point) float32
	SetValue(point image.Point, value float32) bool
	InBounds(point image.Point) bool
	IsObstacle(point image.Point) bool
	Data() [][]float32
}

type field struct {
	width  int
	height int
	data   [][]float32
}

func (f field) Width() int {
	return f.width
}

func (f field) Height() int {
	return f.height
}

// Установить значение в ячейку, если в пределах поля
func (f *field) SetValue(p image.Point, value float32) bool {
	if !f.InBounds(p) {
		return false
	}

	f.data[p.Y][p.X] = value
	return true
}

// Получить значение на поле. Если за пределами, то возвращает препятствие.
func (f field) Value(p image.Point) float32 {
	if !f.InBounds(p) {
		return Obstacle
	}

	return f.data[p.Y][p.X]
}

// В пределах поля
func (f field) InBounds(p image.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < f.width && p.Y < f.height
}

// Точка является препятствием
func (f field) IsObstacle(p image.Point) bool {
	return !f.InBounds(p) || f.Value(p) == Obstacle
}

func (f field) Data() [][]float32 {
	return f.data
}

func NewFieldFromImage(im image.Image, colorAsObstacle bool, colors ...color.Color) IField {
	f := &field{
		width:  im.Bounds().Max.X,
		height: im.Bounds().Max.Y,
		data:   [][]float32{},
	}

	f.data = make([][]float32, f.height)
	for k := range f.data {
		f.data[k] = make([]float32, f.width)
	}

	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			exists := existsColor(im.At(x, y), colors)
			if colorAsObstacle && !exists {
				continue

			} else if !colorAsObstacle && exists {
				continue
			}

			f.SetValue(image.Point{
				X: x,
				Y: y,
			}, Obstacle)
		}
	}

	return f
}

func NewField(data [][]float32) IField {
	f := &field{
		width:  0,
		height: 0,
		data:   data,
	}

	f.width = len(data[0])
	f.height = len(data)
	return f
}

func existsColor(c color.Color, colors []color.Color) bool {
	for _, item := range colors {
		if item == c {
			return true
		}
	}

	return false
}
