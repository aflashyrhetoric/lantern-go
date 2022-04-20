package db

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/models"
)

type Relationship struct {
	*models.Relationship
}

func GetAllRelationships() ([]*models.Relationship, error) {
	points := []*models.Relationship{}
	err := conn.Select(&points, "SELECT * FROM relationships")

	if err != nil {
		return nil, err
	}

	return points, nil
}

func GetRelationshipsForPerson(id string, userID int64) ([]models.RelationshipHydrated, error) {
	relationships := []models.Relationship{}
	err := conn.Select(&relationships, "SELECT id, person_one_id, person_two_id, relationship_type FROM relationships WHERE person_one_id = $1 OR person_two_id = $1", id)
	if err != nil {
		return nil, err
	}

	reorientedWithPersonAsPersonOne := []models.Relationship{}
	relationshipsHydrated := []models.RelationshipHydrated{}
	if len(relationships) > 0 {
		people, err := GetAllPeople(userID)
		if err != nil {
			return nil, err
		}
		// Ensure that any and all relationships are positioned with p1 as the "owner" of the relationship
		for _, relationship := range relationships {
			if id == fmt.Sprint(relationship.PersonOneID) {
				reorientedWithPersonAsPersonOne = append(reorientedWithPersonAsPersonOne, relationship)
			}
			if id == fmt.Sprint(relationship.PersonTwoID) {
				reoriented := models.Relationship{
					ID:               relationship.ID,
					RelationshipType: relationship.RelationshipType,
					PersonOneID:      relationship.PersonTwoID, // SWAP
					PersonTwoID:      relationship.PersonOneID, // SWAP
				}
				reorientedWithPersonAsPersonOne = append(reorientedWithPersonAsPersonOne, reoriented)
			}
		}

		p := people[0]
		for _, relationship := range reorientedWithPersonAsPersonOne {
			r := models.RelationshipHydrated{
				PersonID:         relationship.PersonTwoID,
				RelationshipType: relationship.RelationshipType,
				Person:           *p, // TODO: FIX, USE A GENERIC FIND METHOD TO FIND THE MATCHING PEOPLE ARRAy
			}
			relationshipsHydrated = append(relationshipsHydrated, r)
		}

	}

	return relationshipsHydrated, nil
}

// func GetPressurePointWithID(id string) (*models.PressurePoint, error) {
// 	point := models.PressurePoint{}
// 	err := conn.Get(&point, "SELECT * FROM pressure_points WHERE id = $1", id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &point, nil
// }

func (r Relationship) Validate() (bool, []string) {
	invalidFields := []string{}

	if r.PersonOneID == 0 {
		invalidFields = append(invalidFields, "PersonOneID")
	}

	if r.PersonTwoID == 0 {
		invalidFields = append(invalidFields, "PersonTwoID")
	}

	if r.Relationship == nil {
		invalidFields = append(invalidFields, "Relationship")
	}

	return len(invalidFields) == 0, invalidFields
}

func CreateRelationship(r *Relationship) error {
	valid, fields := r.Validate()

	if !valid {
		return fmt.Errorf("following parameters to the CreateRelationship func was not provided: %v", fields)
	}
	_, err := conn.NamedExec("INSERT into relationships (person_one_id, person_two_id, relationship_type) VALUES (:person_one_id, :person_two_id, :relationship_type)", &r)
	if err != nil {
		return err
	}

	return err
}

// func UpdatePressurePoint(id string, n *models.PressurePoint) error {
// 	_, err := conn.NamedExec("UPDATE pressure_points SET description=:description WHERE id=:id", n)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func Deleterelationship(id string) error {
	_, err := conn.Exec("DELETE FROM relationships WHERE id=$1", id)
	if err != nil {
		return err
	}

	return err
}
