package db

import (
	"github.com/aflashyrhetoric/lantern-go/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*models.User
}

func GetUser(id string) (*models.User, error) {
	user := models.User{}
	err := conn.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := conn.Get(&user, "SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// func GetUserData(id string) (*models.User, error) {
// 	user := models.User{}
// 	err := conn.Get(&user, "SELECT * FROM users WHERE id = $1", id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

func CreateUser(p *User) error {

	userPW := p.Password
	hashed, err := bcrypt.GenerateFromPassword([]byte(userPW), 10)
	if err != nil {
		return err
	}
	p.Password = string(hashed[:])

	_, err = conn.NamedExec("INSERT into users (email, password, created_at) VALUES (:email, :password, :created_at)", &p)
	if err != nil {
		return err
	}

	return err
}

// func UpdateUser(id string, p *models.User) error {
// 	_, err := conn.NamedExec("UPDATE users SET first_name=:first_name, last_name=:last_name, career=:career, mobile=:mobile, email=:email, address=:address, dob=:dob WHERE id=:id", p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func DeleteUser(id string) error {

	tx, err := conn.Begin()

	_, err = tx.Exec("DELETE FROM people WHERE user_id=$1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM notes WHERE user_id=$1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM pressure_points WHERE user_id=$1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}
