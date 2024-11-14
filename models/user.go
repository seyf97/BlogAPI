package models

import (
	"errors"

	"github.com/seyf97/BlogAPI/db"
	"github.com/seyf97/BlogAPI/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) SaveDB() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Hash Password
	hashed_password, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashed_password)
	return err
}

func (u *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var hash_pass_db string

	err := row.Scan(&u.ID, &hash_pass_db)
	if err != nil {
		return err
	}

	is_same := utils.CheckPasswordHash(u.Password, hash_pass_db)
	if !is_same {
		return errors.New("password or email is incorrect")
	}

	return nil

}
