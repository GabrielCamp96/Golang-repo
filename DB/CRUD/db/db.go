package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	var conn string = "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", conn)
	if erro != nil {
		return nil, erro
	}

	erro = db.Ping()
	if erro != nil {
		return nil, erro
	}

	return db, erro
}
