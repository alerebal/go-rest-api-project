package models

import (
	"github.com/alerebal/go-rest-api/db"
	"github.com/alerebal/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	// result, err := stmt.Exec(u.Email, hashedPassword)
	// if err != nil {
	// 	return err
	// }
	// userId, err := result.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	// u.ID = userId

	_, err = stmt.Exec(u.Email, hashedPassword)

	return err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	var users []User
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	user := User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Delete() error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.DB.Exec(query, u.ID)

	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}
}
