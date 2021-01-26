package db

import (
	"fmt"

	"github.com/aflashyrhetoric/lantern-go/models"
)

type Note struct {
	*models.Note
}

func GetAllNotes() ([]*models.Note, error) {
	notes := []*models.Note{}
	err := conn.Select(&notes, "SELECT * FROM Notes")

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func GetNoteWithID(id string) (*models.Note, error) {
	note := models.Note{}
	err := conn.Get(&note, "SELECT * FROM notes WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n Note) Validate() (bool, []string) {
	invalidFields := []string{}

	if n.Text == "" {
		invalidFields = append(invalidFields, "Text")
	}

	return len(invalidFields) == 0, invalidFields
}

func CreateNote(n *Note) error {
	valid, fields := n.Validate()

	if !valid {
		return fmt.Errorf("Following parameters to the CreateNote func was not provided: %v", fields)
	}
	_, err := conn.NamedExec("INSERT into notes (text, person_id) VALUES (:text, :person_id)", &n)
	if err != nil {
		return err
	}

	return err
}

func UpdateNote(id string, n *Note) error {
	_, err := conn.NamedExec("UPDATE notes SET text=:text WHERE id=:id", n)
	if err != nil {
		return err
	}

	return nil
}

func DeleteNote(id string) error {
	_, err := conn.Exec("DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		return err
	}

	return err
}
