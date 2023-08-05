package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(10, 10)
	fmt.Println(calcDistance(*p1, *p2))
}

func calcDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Abs(math.Pow(p1.x, 2) - math.Pow(p2.x, 2) + math.Pow(p1.y, 2) - math.Pow(p2.y, 2)))
}
