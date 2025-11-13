package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (event *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id

	return err
}

// Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).
// This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// if fetching use Query, otherwise use Exec
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err //changing return type to pointer of event allows us to return nil here
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name=?, description=?, location=?, dateTime=?
	WHERE id=?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id=? AND user_id=?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, e.ID)

	return err
}
