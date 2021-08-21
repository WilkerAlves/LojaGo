package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	// db, err := sql.Open("mysql", "root:4874021993@/aluraloja")
	conexao := "root:4874021993@tcp(127.0.0.1:3306)/aluraloja"
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err)
	}
	return db
}
