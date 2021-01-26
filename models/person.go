package models

import "time"

type Person struct {
	ID        int        `db:"id, omitempty" json:"id"`
	FirstName string     `db:"first_name, omitempty" json:"first_name"`
	LastName  string     `db:"last_name, omitempty" json:"last_name"`
	Career    string     `db:"career, omitempty" json:"career"`
	Mobile    string     `db:"mobile, omitempty" json:"mobile"`
	Email     string     `db:"email, omitempty" json:"email"`
	Address   string     `db:"address, omitempty" json:"address"`
	DOB       *time.Time `db:"dob, omitempty" json:"dob"`
}
