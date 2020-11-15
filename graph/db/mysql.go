package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDB *sql.DB

func DbConn() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/graphqlapi")
	if err != nil {
		panic(err.Error())
	}
	MysqlDB = db
}
