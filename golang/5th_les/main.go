package main

import "fmt"

func main() {
	defer handlerPanic()
	fmt.Println("main()")
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	messages[4] = "message 5"

	fmt.Println(messages)
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("Выполнилось успешно")
}
