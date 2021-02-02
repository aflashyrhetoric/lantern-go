package models

type User struct {
	ID       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password"`
}

type UserRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password"`
}
