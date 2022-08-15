package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		message, err := prediction("thu")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(message)
		}
	*/

	// fmt.Println(findMin(1, 2, 3, 4, 5, -55, 3, -2, 44))

	func() {
		fmt.Println("анонимная фукнция")
	}()

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

/*
	Структура switch case
*/

func prediction(day0fWeek string) (string, error) {
	switch day0fWeek {
	case "mon":
		return "Happy start of week", nil
	case "tue":
		return "Happy tuesday", nil
	case "wed":
		return "Happy wednesday", nil
	case "thu":
		return "Happy thusday", nil
	case "fri":
		return "Happy friday", nil
	case "sat":
		return "Happy Saturday", nil
	default:
		return "unknown day", errors.New("invalid day of week")
	}
}

/*
	args and kwargs
*/

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}

/*
	Анонимные функции
*/

/*
	замыкание
*/

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
