package shape

import (
	"math"
)

type Shape interface {
	SapeWithArea
	ShapeWithPerimeter
}

type SapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLenght float32
}

func NewSquare(length float32) Square {
	return Square{
		sideLenght: length,
	}
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (s Square) Perimeter() float32 {
	return s.sideLenght * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
}
