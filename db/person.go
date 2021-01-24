package db

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	ID        int        `db:"id, omitempty"`
	FirstName string     `db:"first_name, omitempty"`
	LastName  string     `db:"last_name, omitempty"`
	Career    string     `db:"career, omitempty"`
	Mobile    string     `db:"mobile, omitempty"`
	Email     string     `db:"email, omitempty"`
	Address   string     `db:"address, omitempty"`
	DOB       *time.Time `db:"dob, omitempty"`
}

func GetAllPeople() ([]*Person, error) {
	people := []*Person{}
	err := conn.Select(&people, "SELECT * FROM people")

	if err != nil {
		return nil, err
	}

	return people, nil
}

func ShowPerson(id string) (*Person, error) {
	person := Person{}
	err := conn.Get(&person, "SELECT * FROM people WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (p Person) Validate() (bool, []string) {
	invalidFields := []string{}

	if p.FirstName == "" {
		invalidFields = append(invalidFields, "FirstName")
	}

	if p.LastName == "" {
		invalidFields = append(invalidFields, "LastName")
	}

	return len(invalidFields) == 0, invalidFields
}

func CreatePerson(p *Person) error {
	valid, fields := p.Validate()

	if !valid {
		spew.Dump(p)
		return fmt.Errorf("Following parameters to the CreatePerson func was not provided: %v", fields)
	}
	_, err := conn.NamedExec("INSERT into people (first_name, last_name, career, mobile, email, address, dob) VALUES (:first_name, :last_name, :career, :mobile, :email, :address, :dob)", &p)
	if err != nil {
		return err
	}

	return err
}

func UpdatePerson(id string, p *Person) error {
	_, err := conn.NamedExec(`UPDATE people SET first_name=:first_name, last_name=:last_name, career=:career, mobile=:mobile, email=:email, address=:address, dob=:dob`, p)
	if err != nil {
		return err
	}

	return err
}

func DeletePerson(id string, p *Person) error {
	_, err := conn.NamedExec(`UPDATE people SET first_name=:first_name, last_name=:last_name, career=:career, mobile=:mobile, email=:email, address=:address, dob=:dob`, p)
	if err != nil {
		return err
	}

	return err
}
