package table

import (
	"github.com/jmoiron/sqlx"
)

var tablesSchema = []string{PositionTableSchema}

func LoadTables(db *sqlx.DB) {
	for _, schema := range tablesSchema {
		db.MustExec(schema)
	}
}
