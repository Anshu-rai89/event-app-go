package models

import (
	"github.com/Anshu-rai89/event-app-go/db"
)

func (e *Event) RegisterEvent(userId int64) error {
	query := `INSERT INTO eventRegistrations (userId, eventId) VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) CancelEventRegistration(userId int64) error {
	query := `DELETE * FROM eventRegistrations WHERE eventId = ? AND userID = ? LIMIT 1`

	smtp, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(e.Id, userId)

	return err
}
