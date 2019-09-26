package wave

import (
	"image"
	"math"
)

/*
Получаем эффективный путь.

field - поле с препятствиями.

startPoint - стартовая точка.

startPoint - финишная точка.
*/
func Search(field IField, startPoint image.Point, finishPoint image.Point) (*Path, bool) {
	field.SetValue(startPoint, 1)
	found := wavePropagation(field, finishPoint, []image.Point{startPoint})
	if !found {
		return nil, false
	}

	path := Path{finishPoint}
	nextPoint(field, &path, startPoint, finishPoint)
	path.Reverse()
	return &path, true
}

func nextPoint(f IField, w *Path, startPoint image.Point, checkPoint image.Point) {
	if startPoint.Eq(checkPoint) {
		return
	}

	newPoint := image.Point{
		X: -1,
		Y: -1,
	}
	//проверяем окружные 8 клеток
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {

			if x == 0 && y == 0 { // checkPoint не рассматриваем
				continue
			}

			p := image.Point{
				X: checkPoint.X + x,
				Y: checkPoint.Y + y,
			}
			if f.IsObstacle(p) || f.Value(p) == Empty || f.Value(p) != Empty && f.Value(checkPoint) < f.Value(p) {
				continue
			}

			if isDiagonal(p, checkPoint) {
				if isDiagonalObstacle(f, p, checkPoint) { // есть диагональные преграды
					continue
				}
			}

			if (newPoint.X == -1 && newPoint.Y == -1) || f.Value(p) < f.Value(newPoint) {
				newPoint = p
			}
		}
	}

	if newPoint.X == -1 && newPoint.Y == -1 {
		newPoint = startPoint
	}

	*w = append(*w, newPoint)
	nextPoint(f, w, startPoint, newPoint)
}

/*
Распространение волны по field

field - поле с препятствиями.

finishPoint - финишная точка

points - список точек, из которых необходимо продолжить волну
*/
func wavePropagation(field IField, finishPoint image.Point, points []image.Point) bool {
	numPoints := len(points)
	if numPoints == 0 {
		return false
	}

	var nextPoints []image.Point
	nextPoints = []image.Point{}

	//обходим точки
	for _, currPoint := range points {

		if currPoint.Eq(finishPoint) {
			return true
		}
		v := field.Value(currPoint) + 1

		//проверяем окружные 8 клеток
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {

				if x == 0 && y == 0 { // currPoint не рассматриваем
					continue
				}

				p := currPoint
				p.X += x
				p.Y += y

				if field.Value(p) != Empty { // ячейка уже обработана/занята/за гранью
					continue
				}

				if isDiagonal(p, currPoint) {
					continue
				}

				if field.Value(p) == Empty {
					nextPoints = append(nextPoints, p)
				}

				if field.Value(p) == Empty || field.Value(p) > v {
					field.SetValue(p, v)
				}
			}
		}
	}

	//повторяем для следующих клеток
	return wavePropagation(field, finishPoint, nextPoints)
}

// две точки создают диагональ?
func isDiagonal(p1 image.Point, p2 image.Point) bool {
	return math.Abs(float64(p1.X-p2.X))+math.Abs(float64(p1.Y-p2.Y)) == 2
}

/*
Есть ли препятствия для диагональной линии

есть препятствие для диагональной линии:

* | ⬛

⬜ | *

и здесь тоже:

* | ⬜

⬛ | *

нет препятствий:

* | ⬜

⬜ | *
*/
func isDiagonalObstacle(f IField, p1 image.Point, p2 image.Point) bool {
	pVert := image.Point{
		X: p1.X,
		Y: p2.Y,
	}
	pHoriz := image.Point{
		X: p2.X,
		Y: p1.Y,
	}
	return f.IsObstacle(pVert) && f.IsObstacle(pHoriz)
}
