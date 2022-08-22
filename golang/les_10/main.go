package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Slatt4561@/vk")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// result, err := db.Exec("insert into vk.users (firstname, lastname, email, password_hash, phone) values (?, ?, ?, ?, ?)",
	//	"Slava", "Galinov", "g4linoff@gmail.com", "qwerty", 9994451337)
	result, err := db.Exec("insert into vk.users (firstname, lastname, email, password_hash, phone) values (?, ?, ?, ?, ?)",
		"Katya", "Shelby", "gate2@gmail.com", "qwerty", 79016660225)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
