package models

import "time"

type Person struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Career    string    `db:"career"`
	Mobile    string    `db:"mobile"`
	Email     string    `db:"email"`
	Address   string    `db:"address"`
	DOB       time.Time `db:"dob"`
}
