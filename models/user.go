package models

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type UserRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password"`
}

// Data that most pages should have
type UserData struct {
	People []*Person `json:"people"`
}
