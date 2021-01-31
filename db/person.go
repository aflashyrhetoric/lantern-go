package db

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/models"
)

type Person struct {
	*models.Person
}

func GetAllPeople() ([]*models.Person, error) {
	people := []*models.Person{}
	err := conn.Select(&people, "SELECT * FROM people")

	if err != nil {
		return nil, err
	}

	return people, nil
}

func GetPerson(id string) (*models.Person, error) {
	person := models.Person{}
	// err := conn.Get(&person, `
	// 	SELECT
	// 		p.id,
	// 		p.first_name,
	// 		p.last_name,
	// 		p.career,
	// 		p.mobile,
	// 		p.email,
	// 		p.address,
	// 		p.dob
	// 	FROM people p
	// 	INNER JOIN notes n
	// 		on n.person_id = $1
	// 	INNER JOIN pressure_points pp
	// 		on pp.person_id = $1
	// 	WHERE
	// 		p.id = $1
	// `, id)
	err := conn.Get(&person, "SELECT * FROM people WHERE id = $1", id)

	notes := []models.Note{}
	err = conn.Select(&notes, "SELECT id, text FROM notes WHERE person_id = $1", id)

	points := []models.PressurePoint{}
	err = conn.Select(&points, "SELECT id, description FROM pressure_points WHERE person_id = $1", id)

	person.Notes = notes
	person.PressurePoints = points

	if err != nil {
		return nil, err
	}

	return &person, nil

}

func GetPersonalData(id string) (*models.Person, error) {
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
		return fmt.Errorf("Following parameters to the CreatePerson func was not provided: %v", fields)
	}
	_, err := conn.NamedExec("INSERT into people (first_name, last_name, career, mobile, email, address, dob) VALUES (:first_name, :last_name, :career, :mobile, :email, :address, :dob)", &p)
	if err != nil {
		return err
	}

	return err
}

func UpdatePerson(id string, p *models.Person) error {
	_, err := conn.NamedExec("UPDATE people SET first_name=:first_name, last_name=:last_name, career=:career, mobile=:mobile, email=:email, address=:address, dob=:dob WHERE id=:id", p)
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
