package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.height, u.sex, u.sex)
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vasya", "Male", 23, 75, 185)
	user2 := User{"Vasya", 23, "Male", 84, 195}

	fmt.Println(user1.name, user1.age)

	user1.printUserInfo()
	user2.printUserInfo()
}
