package main

import (
	"errors"
	"fmt"
)

var msg string

/*
init - перед main. Для иницилизации pocket.
*/
func init() {
	msg = "from init()"
}

func main() {
	// fmt.Println(msg)
	message := "The best wish"
	fmt.Println(message)
	changeMessage(&message)
	fmt.Println(message)

	var p *int
	number := 5
	p = &number
	fmt.Println(p, &number)
	*p = 10
	fmt.Println(number)

	/*
		Массив
	*/

	// messages := [3]string{"1", "2", "3"}

	messages := []string{"1", "2", "3"}
	fmt.Println(messages)
	printMessages(messages)
}

/*
	Ссылки и указатели
*/

func changeMessage(message *string) {
	*message += " (from func changeMessage())"
}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("emtpy array")
	}

	messages[1] = "5"

	fmt.Println(messages)

	return nil
}
