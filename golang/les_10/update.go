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
	// обновляем строку с id=1
	result, err := db.Exec("update vk.users set phone = ? where id = ?", 9994450000, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
