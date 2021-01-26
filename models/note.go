package models

type Note struct {
	ID       int    `db:"id" json:"id"`
	Text     string `db:"text" json:"text"`
	PersonID int    `db:"person_id" json:"person_id"`
}
