package db

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// This time the global variable is unexported.
var db *sqlx.DB

// InitDB sets up setting up the connection pool global variable.
func Start(dataSourceName string) error {
	var err error

	db, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return db.Ping()
}
