package db

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Career    string    `json:"career" db:"career"`
	Mobile    string    `json:"mobile" db:"mobile"`
	Email     string    `json:"email" db:"email"`
	Address   string    `json:"address" db:"address"`
	DOB       time.Time `json:"dob" db:"dob"`
}

func GetAllPeople() ([]*Person, error) {
	people := []*Person{}
	db.Select(&people, "SELECT * FROM people")

	spew.Dump(people)

	return people, nil
}
