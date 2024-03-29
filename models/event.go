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
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
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

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `UPDATE events SET name= ?, description=?, location=?, dateTime=?
		WHERE id = ?`

	smtp, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(&event.Name, &event.Description, &event.Location, &event.DateTime, &event.Id)

	return err

}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`

	smtp, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(id)

	return err

}
