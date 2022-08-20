package main

import (
	"fmt"
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
	square := Square{5}
	//circle := Circle{8}

	printShapeArea(square)
	//printShapeArea(circle)

	printInterface(22)
	printInterface("qwe")
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func printInterface(i interface{}) {
	/*
		switch t := i.(type) {
		case int:
			fmt.Println("int", t)
		case bool:
			fmt.Println("bool", t)
		default:
			fmt.Println("Unknown type", t)
		}
		fmt.Printf("%+v\n", i)
	*/
	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not streing")
		return
	}
	fmt.Println(len(str))
}
