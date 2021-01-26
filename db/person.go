package db

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/models"
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	*models.Person
	// ID        int        `db:"id, omitempty" json:"id"`
	// FirstName string     `db:"first_name, omitempty" json:"first_name"`
	// LastName  string     `db:"last_name, omitempty" json:"last_name"`
	// Career    string     `db:"career, omitempty" json:"career"`
	// Mobile    string     `db:"mobile, omitempty" json:"mobile"`
	// Email     string     `db:"email, omitempty" json:"email"`
	// Address   string     `db:"address, omitempty" json:"address"`
	// DOB *time.Time `db:"dob, omitempty" json:"dob"`
}

func GetAllPeople() ([]*models.Person, error) {
	people := []*models.Person{}
	err := conn.Select(&people, "SELECT * FROM people")

	if err != nil {
		return nil, err
	}

	return people, nil
}

func GetPersonWithID(id string) (*models.Person, error) {
	person := models.Person{}
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

func UpdatePerson(id string, p *models.Person) error {
	spew.Dump(p)

	_, err := conn.NamedExec("UPDATE people SET first_name=:first_name, last_name=:last_name, career=:career, mobile=:mobile, address=:address, dob=:dob WHERE id=:id", p)
	if err != nil {
		return err
	}

	return nil
}

func DeletePerson(id string) error {
	_, err := conn.Exec("DELETE FROM people WHERE id=$1", id)
	if err != nil {
		return err
	}

	return err
}
