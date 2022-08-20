package main

import (
	"fmt"
	"golang_ninja/basic/shape"

	"github.com/zhashkevych/scheduler"
)

func main() {
	square := shape.NewSquare(5)
	//circle := Circle{8}

	printShapeArea(square)
	//printShapeArea(circle)

	scheduler := scheduler.NewScheduler()

	printInterface(22)
	printInterface("qwe")
}

func printShapeArea(s shape.Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
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
