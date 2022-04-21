package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	Server().Run()
	//fmt.Println("test")
	//DBHOST := os.Getenv("MYSQL_DBNAME")
	//fmt.Println(DBHOST)
}
