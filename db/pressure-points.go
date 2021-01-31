package db

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/models"
)

type PressurePoint struct {
	*models.PressurePoint
}

func GetAllPressurePoints() ([]*models.PressurePoint, error) {
	points := []*models.PressurePoint{}
	err := conn.Select(&points, "SELECT * FROM pressure_points")

	if err != nil {
		return nil, err
	}

	return points, nil
}

func GetPressurePointsForPerson(id string) ([]models.PressurePoint, error) {
	points := []models.PressurePoint{}
	err := conn.Select(&points, "SELECT id, description FROM pressure_points WHERE person_id = $1", id)
	if err != nil {
		return nil, err
	}

	return points, nil
}

// func GetPressurePointWithID(id string) (*models.PressurePoint, error) {
// 	point := models.PressurePoint{}
// 	err := conn.Get(&point, "SELECT * FROM pressure_points WHERE id = $1", id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &point, nil
// }

func (n PressurePoint) Validate() (bool, []string) {
	invalidFields := []string{}

	if n.Description == "" {
		invalidFields = append(invalidFields, "Description")
	}

	return len(invalidFields) == 0, invalidFields
}

func CreatePressurePoint(n *PressurePoint) error {
	valid, fields := n.Validate()

	if !valid {
		return fmt.Errorf("Following parameters to the CreatePressurePoint func was not provided: %v", fields)
	}
	_, err := conn.NamedExec("INSERT into pressure_points (description, person_id) VALUES (:description, :person_id)", &n)
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

func DeletePressurePoint(id string) error {
	_, err := conn.Exec("DELETE FROM pressure_points WHERE id=$1", id)
	if err != nil {
		return err
	}

	return err
}
