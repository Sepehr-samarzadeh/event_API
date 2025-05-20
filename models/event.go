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

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // use Exec when we  want to add something (change to db) && use query when you want to fetch it

	if err != nil {
		return nil, err
	}
	defer rows.Close() //make sure its close when accessing it

	var events []Event

	for rows.Next() { //next will keep the loop running as long as  we have rows to read
		var even Event
		err := rows.Scan(&even.ID, &even.Name, &even.Description, &even.Location, &even.DateTime, &even.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, even)

	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ? "
	row := db.DB.QueryRow(query, id) //since we only need the row with specific id not bunch of rows
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
