package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	urlConexao := "root:root@tcp(0.0.0.0:3306)/godatabase?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConexao)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão database aberta")
}