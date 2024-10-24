package models

import (
	"errors"
	"go_rest_api/db"
	"time"
)

type Event struct {
	ID          int    `json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Location    string `binding:"required" json:"location"`
	DateTime    int64  `binding:"required" json:"dateTime"` // Store epoch time as int64
	UserId      int    `binding:"required" json:"userId"`
}


func (e *Event) Save() error {

	query := `INSERT INTO events (name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DBClient.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	// Execute the query
	result , err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	// Set the ID field of the struct
	e.ID = int(id)

	return nil
}

func GetAllEvents() ([]Event,error){
	
	rows, err := db.DBClient.Query("SELECT * FROM events")

	if err != nil {	
		return nil, err
	}

	// Close the rows after the function ends
	defer rows.Close()

	events := []Event{}

	// Iterate over the rows
	for rows.Next() {
		var e Event

		// Scan the row into the struct
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}

		// Append the struct to the slice
		events = append(events, e)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	var e Event

	row := db.DBClient.QueryRow("SELECT * FROM events WHERE id = ?", id)
	
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

	if err != nil {
		// Return nil and the custom error
		return nil, errors.New("Event not found")

	}

	return &e, nil
}

// Helper function to convert epoch to time.Time
func EpochToTime(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

// Helper function to convert time.Time to epoch
func TimeToEpoch(t time.Time) int64 {
	return t.Unix()
}