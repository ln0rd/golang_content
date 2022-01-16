package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	urlConexao := "root:root@tcp(0.0.0.0:3306)/godatabase?charset=utf8&parseTime=True&loc=Local"

	fmt.Println("Connecting database")
	db, err := sql.Open("mysql", urlConexao)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	
	return db, nil
}