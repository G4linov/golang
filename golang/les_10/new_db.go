package main

import "database/sql"

func main() {
	db, err := sql.Open("mysql", "root:Slatt4561@")
	if err != nil {
		panic(err)
	}
	db.Close()
}
