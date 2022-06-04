package main

import (
	"fmt"
	"math"
)

// point представляет собой объект точки, хранящие в себе ее координаты
type point struct {
	x, y float64
}

// NewPoint - конструктор point
func NewPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

// line представляет собой объект линии, хранящие в себе две точки, формирующие данную линию
type line struct {
	fpoint *point
	spoint *point
}

// NewLine - конструктор line
func NewLine(first, second *point) *line {
	return &line{
		fpoint: first,
		spoint: second,
	}
}

// Length возвращает длину линии l по теореме пифагора
func (l *line) Length() float64 {
	pythagoras := math.Pow(l.spoint.x-l.fpoint.x, 2) + math.Pow(l.spoint.y-l.fpoint.y, 2)
	return math.Sqrt(pythagoras)
}

func main() {
	line := NewLine(NewPoint(5, 10), NewPoint(10, 5))
	fmt.Println(line.Length())
}
