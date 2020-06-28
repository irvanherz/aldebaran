package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB = nil

func SetupDB() {
	db, err := sql.Open("mysql", "root"+":"+""+"@/"+"aldebaran")
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	} else {
		Db = db
	}
}
