package database

import (
	"bispy-agent/database/table"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func Connect() {

	cwd, getwdErr := os.Getwd()
	if getwdErr != nil {
		log.Fatal(getwdErr)
	}

	db, connectErr := sqlx.Connect("sqlite3", cwd+os.Getenv("BISPY_DB_PATH"))
	if connectErr != nil {
		log.Fatal(connectErr)
	}

	table.LoadTables(db)
	DB = db
}
