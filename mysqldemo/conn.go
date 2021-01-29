package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := "kumit:qaz78963@tcp(35.201.213.247)/learngo"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mysql server conn!")
}
