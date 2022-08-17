package main

import (
	"fmt"
	"reflect"
)

var e, r, t int

func main() {
	message_my := ("Я скоро стану GOLANG ninja")
	// var message int
	// message = 12
	fmt.Println(reflect.TypeOf(message_my))

	/*
		Типы перменных
	*/
	var message string
	var number int
	var crazy float32
	var test bool
	// var a byte = 62
	// var b rune = 'g'
	crazy = 12.2
	test = true
	smin := []byte("sad")

	fmt.Println(message, number, crazy, test, smin)
	// fmt.Println(b)

	/*
		Множественное присвоение
	*/

	a, b, c := 1, 2, 3
	a, b = b, a
	fmt.Println(a, b, c)

	/*
		Видимость перменных
	*/

	e, r, t = 3, 2, 1
	print()
}

func print() {
	fmt.Println(e, r, t)
}
