package models

import (
	"time"
)

type Person struct {
	ID        int        `db:"id, omitempty" json:"id"`
	FirstName string     `db:"first_name" json:"first_name"`
	LastName  string     `db:"last_name" json:"last_name"`
	Career    string     `db:"career" json:"career"`
	Mobile    string     `db:"mobile" json:"mobile"`
	Email     string     `db:"email" json:"email"`
	Address   string     `db:"address" json:"address"`
	DOB       *time.Time `db:"dob" json:"dob"`

	Notes          []Note          `json:"notes,omitempty"`
	PressurePoints []PressurePoint `json:"pressure_points,omitempty"`
}

type PersonRequest struct {
	ID        int     `db:"id, omitempty" json:"id"`
	FirstName string  `db:"first_name" json:"first_name"`
	LastName  string  `db:"last_name" json:"last_name"`
	Career    string  `db:"career" json:"career"`
	Mobile    string  `db:"mobile" json:"mobile"`
	Email     string  `db:"email" json:"email"`
	Address   string  `db:"address" json:"address"`
	DOB       *string `db:"dob" json:"dob"`

	Notes          []Note          `json:"notes,omitempty"`
	PressurePoints []PressurePoint `json:"pressure_points,omitempty"`
}
