package models

import (
	"time"

	"github.com/Anshu-rai89/event-app-go/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	 INSERT INTO events (name , description, location, dateTime, user_id) 
	 VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvents() []Event {
	return events
}