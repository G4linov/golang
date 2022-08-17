package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 48,
	}

	users["Serega"] = 42
	delete(users, "Serega")

	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println(len(users))
}
