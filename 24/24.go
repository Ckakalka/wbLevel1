package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(first Point, second Point) float64 {
	deltaX := first.x - second.x
	deltaY := first.y - second.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func main() {
	a := NewPoint(2, 3)
	b := NewPoint(-1, 7)
	fmt.Println(Distance(a, b))
}
