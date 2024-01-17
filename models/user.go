package models

import "github.com/Anshu-rai89/event-app-go/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(name, email, password) VALUES (?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	results, err := stmt.Exec(u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := results.LastInsertId()

	u.ID = userId
	return err
}
