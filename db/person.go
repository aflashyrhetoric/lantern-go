package db

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

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

func GetAllPeople() ([]*Person, error) {
	people := []*Person{}
	err := conn.Select(&people, "SELECT * FROM people")

	if err != nil {
		return nil, err
	}

	return people, nil
}

func CreatePerson(p Person) error {
	spew.Dump(conn)
	_, err := conn.NamedExec("INSERT into people (first_name, last_name, career, mobile, email, address, dob) VALUES (:first_name, :last_name, :career, :mobile, :email, :address, :dob)", &p)
	if err != nil {
		return err
	}

	return err
}
