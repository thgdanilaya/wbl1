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
	return &Point{x: x, y: y}
}

func (p *Point) DistanceTo(other *Point) float64 {
	return math.Hypot(p.x-other.x, p.y-other.y)
}

func main() {
	a := NewPoint(3, 4)
	b := NewPoint(5, 6)
	fmt.Println(a.DistanceTo(b))
}
