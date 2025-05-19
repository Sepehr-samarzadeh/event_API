package models

import (
	"time"

	"sep.com/eventapi/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(name,description,location,dateTime,user_id) VALUES(
	?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // no matter we get error or not close this statement
	// in other words make sure you close it after executing the command
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId() //to get the last event id which inserted

	e.ID = id
	return err

}

func GetAllEvents() []Event {
	return events
}
