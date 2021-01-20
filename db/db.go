package db

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// This time the global variable is unexported.
var conn *sqlx.DB

// Start ... sets up setting up the connection pool global variable.
func Start(dataSourceName string) error {
	var err error

	conn, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return conn.Ping()

}
