package main

import (
	"fmt"
	"strings"
)

// ПРОБЛЕМА: Имеем векторное изображение без каких-либо методов
// Нужно нарисовать данное изображение в растровом формате, не добавляя/модифицируя
// методы VectorImage

// Line представляет из себе координаты точек начала и конца линии, которые ее формируют
type Line struct {
	X1, Y1, X2, Y2 int
}

// VectorImage - слайс линий, который формирует изображение
type VectorImage struct {
	Lines []Line
}

// NewRectangle - возвращает векторную картинку, представляющий из себе прямоугольник
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{width, 0, width, height},
			{width, height, 0, height},
			{0, height, 0, 0},
		},
	}
}

// Point хранит в себе координаты точки
type Point struct {
	X, Y int
}

// RasterImage - интерфейс, описывающий методы растрового изображения
type RasterImage interface {
	GetPoints() []Point
}

// DrawPoints формирует "растровое" представление картинки
func DrawPoints(img RasterImage) string {
	maxX, maxY := 0, 0
	points := img.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}

		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}

	maxX++
	maxY++

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// vectorToRasterAdapter - адаптер, хранящий в себе массив точек
type vectorToRasterAdapter struct {
	points []Point
}

// addLine возвращает растровое представление линии векторного изображения
func (v *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)

	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
}

// GetPoints возвращает количество точек
func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

// VectorToRaster адаптирует векторное изображение к растровому
func VectorToRaster(v *VectorImage) RasterImage {
	// Инициализируем адаптер, который удовлетворяет интерфейсу RasterImage
	adapter := vectorToRasterAdapter{}

	// Формируем картинку
	for _, line := range v.Lines {
		adapter.addLine(line)
	}

	return &adapter
}

func main() {
	rc := NewRectangle(6, 4)

	// Адаптируемся
	a := VectorToRaster(rc)

	// ОПА, РИСУЕМ
	fmt.Println(DrawPoints(a))
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}
