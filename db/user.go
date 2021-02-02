package db

import (
	"github.com/aflashyrhetoric/lantern-go/models"
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

// func GetUserData(id string) (*models.User, error) {
// 	user := models.User{}
// 	err := conn.Get(&user, "SELECT * FROM users WHERE id = $1", id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

func CreateUser(p *User) error {
	_, err := conn.NamedExec("INSERT into users (email, password) VALUES (:email, :password)", &p)
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
