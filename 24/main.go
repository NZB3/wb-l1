package main

import "math"

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

func GetDistance(p1, p2 Point) float64 {
	x1 := p1.GetX()
	y1 := p1.GetY()

	x2 := p2.GetX()
	y2 := p2.GetY()

	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(2, 2)

	println(GetDistance(*p1, *p2))
}
