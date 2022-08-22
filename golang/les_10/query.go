package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id            int
	firstname     string
	lastname      string
	email         string
	password_hash string
	phone         int
}

func main() {
	db, err := sql.Open("mysql", "root:Slatt4561@/vk")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from vk.users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []user{}

	for rows.Next() {
		p := user{}
		err := rows.Scan(&p.id, &p.firstname, &p.lastname, &p.email, &p.password_hash, &p.phone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	for _, p := range users {
		fmt.Println(p.id, p.firstname, p.lastname, p.email, p.password_hash, p.phone)
	}
}
