package wave

import "image"

type Path []image.Point

func (p Path) Reverse() {
	for i := len(p)/2 - 1; i >= 0; i-- {
		opp := len(p) - 1 - i
		p[i], p[opp] = p[opp], p[i]
	}
}

func (p Path) Contains(point image.Point) bool {
	for _, item := range p {
		if item.Eq(point) {
			return true
		}
	}
	return false
}

func (p Path) CommonPoints(path Path) []image.Point {
	var result []image.Point
	result = []image.Point{}

	for _, item := range path {
		if p.Contains(item) {
			result = append(result, item)
		}
	}

	return result
}
