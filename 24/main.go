package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

type line struct {
	fpoint *point
	spoint *point
}

func main() {
	line := NewLine(NewPoint(5, 10), NewPoint(10, 5))
	fmt.Println(line.Length())
}

func NewPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func NewLine(first, second *point) *line {
	return &line{
		fpoint: first,
		spoint: second,
	}
}

func (l *line) Length() float64 {
	pythagoras := math.Pow(l.spoint.x-l.fpoint.x, 2) + math.Pow(l.spoint.y-l.fpoint.y, 2)
	return math.Sqrt(pythagoras)
}
