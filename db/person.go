package db

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	ID        int       `db:"id, omitempty"`
	FirstName string    `db:"first_name, omitempty"`
	LastName  string    `db:"last_name, omitempty"`
	Career    string    `db:"career, omitempty"`
	Mobile    string    `db:"mobile, omitempty"`
	Email     string    `db:"email, omitempty"`
	Address   string    `db:"address, omitempty"`
	DOB       time.Time `db:"dob, omitempty"`
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

func UpdatePerson(id int, p Person) error {

	_, err := conn.NamedExec(`
			UPDATE people 
			SET 
				first_name=:first_name,
				last_name=:last_name,
				career=:career,
				mobile=:mobile,
				email=:email,
				address=:address,
				dob=:dob,
			`, &p)
	if err != nil {
		return err
	}

	return err
}
