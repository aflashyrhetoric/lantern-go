package db

import (
	"time"
)

type Person struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name" faker:"first_name"`
	LastName  string    `json:"last_name" db:"last_name" faker:"last_name"`
	Career    string    `json:"career" db:"career" faker:"word"`
	Mobile    string    `json:"mobile" db:"mobile" faker:"oneof: doctor, coder, engineer, chef"`
	Email     string    `json:"email" db:"email" faker:"email"`
	Address   string    `json:"address" db:"address" faker:"word"`
	DOB       time.Time `json:"dob" db:"dob" faker:"date"`
}

func GetAllPeople() ([]*Person, error) {
	people := []*Person{}
	err := db.Select(&people, "SELECT * FROM people")

	if err != nil {
		return nil, err
	}
	return people, nil
}

func CreatePerson(p Person) error {
	_, err := db.Exec("INSERT into people (first_name, last_name, career, mobile, email, address, dob) VALUES ($1, $2, $3, $4, $5, $6, $7)", p.FirstName, p.LastName, p.Career, p.Mobile, p.Email, p.Address)
	if err != nil {
		return err
	}

	return err
}
