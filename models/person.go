package models

import (
	"time"
)

type Person struct {
	ID        int        `db:"id, omitempty" json:"id,omitempty"`
	FirstName string     `db:"first_name" json:"first_name,omitempty"`
	LastName  string     `db:"last_name" json:"last_name,omitempty"`
	Career    string     `db:"career" json:"career,omitempty"`
	Mobile    string     `db:"mobile" json:"mobile,omitempty"`
	Email     string     `db:"email" json:"email,omitempty"`
	Address   string     `db:"address" json:"address,omitempty"`
	DOB       *time.Time `db:"dob" json:"dob,omitempty"`
	UserID    int64      `db:"user_id" json:"user_id,omitempty"`

	RelationshipToUser              RelationshipType `db:"" json:"relationship_to_user"`
	RelationshipToUserThroughPerson int              `db:"" json:"relationship_to_user_through_person"`

	Notes          []Note                 `json:"notes"`
	PressurePoints []PressurePoint        `json:"pressure_points"`
	Relationships  []RelationshipHydrated `json:"relationships"`
}

type CreatePersonRequest struct {
	FirstName                       string           `db:"first_name" json:"first_name"`
	LastName                        string           `db:"last_name" json:"last_name"`
	Career                          string           `db:"career" json:"career"`
	Mobile                          string           `db:"mobile" json:"mobile"`
	Email                           string           `db:"email" json:"email"`
	Address                         string           `db:"address" json:"address"`
	DOB                             *string          `db:"dob" json:"dob"`
	RelationshipToUser              RelationshipType `db:"" json:"relationship_to_user"`
	RelationshipToUserThroughPerson int              `db:"" json:"relationship_to_user_through_person"`
}

type UpdatePersonRequest struct {
	FirstName                       string           `db:"first_name" json:"first_name"`
	LastName                        string           `db:"last_name" json:"last_name"`
	Career                          string           `db:"career" json:"career"`
	Mobile                          string           `db:"mobile" json:"mobile"`
	Email                           string           `db:"email" json:"email"`
	Address                         string           `db:"address" json:"address"`
	DOB                             *string          `db:"dob" json:"dob"`
	RelationshipToUser              RelationshipType `db:"" json:"relationship_to_user"`
	RelationshipToUserThroughPerson int              `db:"" json:"relationship_to_user_through_person"`
}
