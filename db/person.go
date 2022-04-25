package db

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/models"
)

type Person struct {
	*models.Person
}

func GetAllPeople(id int64) ([]*models.Person, error) {
	people := []*models.Person{}
	err := conn.Select(&people, "SELECT * FROM people where user_id = $1", id)

	if err != nil {
		return nil, err
	}

	return people, nil
}

func GetPerson(id string, userID int64) (*models.Person, error) {
	person := models.Person{}
	err := conn.Get(&person, "SELECT * FROM people WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	// notes := []models.Note{}
	// err = conn.Select(&notes, "SELECT id, text FROM notes WHERE person_id = $1", id)
	notes, err := GetNotesForPerson(id)
	if err != nil {
		return nil, err
	}

	// points := []models.PressurePoint{}
	// err = conn.Select(&points, "SELECT id, description FROM pressure_points WHERE person_id = $1", id)
	points, err := GetPressurePointsForPerson(id)
	if err != nil {
		return nil, err
	}

	relationships, err := GetRelationshipsForPerson(id, userID)
	if err != nil {
		return nil, err
	}

	person.Notes = notes
	person.PressurePoints = points
	person.Relationships = relationships

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
	valid, invalidFields := p.Validate()

	if !valid {
		return fmt.Errorf("following parameters to the CreatePerson func was not provided: %v", invalidFields)
	}

	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	// Save the person, returning ID
	_, err = tx.NamedExec("INSERT into people (first_name, last_name, career, mobile, email, address, dob, relationship_to_user, relationship_to_user_through_person_id, user_id) VALUES (:first_name, :last_name, :career, :mobile, :email, :address, :dob, :relationship_to_user, :relationship_to_user_through_person_id, :user_id) returning id", &p)
	if err != nil {
		return err
	}

	// newlyInsertedID, err := result.LastInsertId()
	// if err != nil {
	// 	return err
	// }

	// Save a relationship between the user and the person
	// _, err = tx.Exec("INSERT into relationships (person_one_id, person_two_id, relationship_type) VALUES ($1, $2, $3)", &p.UserID, &newlyInsertedID, &p.RelationshipToUser)
	// if err != nil {
	// 	return err
	// }

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdatePerson(id string, p *models.Person) error {
	_, err := conn.NamedExec("UPDATE people SET first_name=:first_name, last_name=:last_name, career=:career, mobile=:mobile, email=:email, address=:address, dob=:dob, relationship_to_user=:relationship_to_user, relationship_to_user_through_person_id=:relationship_to_user_through_person_id WHERE id=:id", p)
	if err != nil {
		return err
	}

	return nil
}

func DeletePerson(id string) error {

	tx, err := conn.Begin()
	_, err = tx.Exec("DELETE FROM notes WHERE person_id=$1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM pressure_points WHERE person_id=$1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM people WHERE id=$1", id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}
