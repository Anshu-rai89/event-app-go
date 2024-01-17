package models

import (
	"errors"

	"github.com/Anshu-rai89/event-app-go/db"
	"github.com/Anshu-rai89/event-app-go/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(name, email, password) VALUES (?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	//hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	results, err := stmt.Exec(u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := results.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidatePassword() error {
	query := `SELECT password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var password string
	err := row.Scan(&password)

	println("Password", password, " ", u.Password)
	if err != nil {
		return errors.New("Invalid creds")
	}
	err = utils.CheckPassword(u.Password, password)
	return err
}
