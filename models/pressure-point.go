package models

type PressurePoint struct {
	ID          int    `db:"id" json:"id"`
	Description string `db:"description" json:"description"`
	PersonID    int    `db:"person_id" json:"person_id"`
}
