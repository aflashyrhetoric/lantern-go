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

	RelationshipToUser              NullString `db:"relationship_to_user" json:"relationship_to_user,omitempty"`
	RelationshipToUserThroughPerson NullInt64  `db:"relationship_to_user_through_person_id" json:"relationship_to_user_through_person_id,omitempty"`

	Notes          []Note                 `json:"notes"`
	PressurePoints []PressurePoint        `json:"pressure_points"`
	Relationships  []RelationshipHydrated `json:"relationships"`
}

type CreatePersonRequest struct {
	FirstName                       string     `json:"first_name"`
	LastName                        string     `json:"last_name"`
	Career                          string     `json:"career"`
	Mobile                          string     `json:"mobile"`
	Email                           string     `json:"email"`
	Address                         string     `json:"address"`
	DOB                             *string    `json:"dob"`
	RelationshipToUser              NullString `json:"relationship_to_user,omitempty"`
	RelationshipToUserThroughPerson NullInt64  `json:"relationship_to_user_through_person_id,omitempty"`
}

type UpdatePersonRequest struct {
	FirstName                       string     `json:"first_name"`
	LastName                        string     `json:"last_name"`
	Career                          string     `json:"career"`
	Mobile                          string     `json:"mobile"`
	Email                           string     `json:"email"`
	Address                         string     `json:"address"`
	DOB                             *string    `json:"dob"`
	RelationshipToUser              NullString `json:"relationship_to_user,omitempty"`
	RelationshipToUserThroughPerson NullInt64  `json:"relationship_to_user_through_person_id,omitempty"`
}
