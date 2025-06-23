package models

import (
	"event-rest-api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events []Event = []Event{}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, date_time, user_id)
		VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	e.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents() []Event {
	return events
}
