package models

import (
	"log/slog"
	"time"

	"go-lang-restapi/db"
)

type Event struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"  binding:"required,min=3"`
	Description string    `json:"description"  binding:"required,min=3"`
	Location    string    `json:"location"  binding:"required"`
	DateTime    time.Time `json:"datetime"  binding:"required"`
	UserId      *uint64   `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	/* Here WE Will Add the Db Code */

	query := "insert into events(title,description,location,dateTime,user_id) values(?,?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		slog.Warn(err.Error())
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	e.ID = uint64(id)
	slog.Info("new Event Added here")
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		slog.Error(err.Error())
		return []Event{}, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
		events = append(events, event)
	}
	slog.Info("fetcging all the events", "events", events)
	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := "select * from events where id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return &event, nil
}

func (e Event) UpdateEvent() error {
	query := "update  events set title = ?,description= ?,location= ?,dateTime=?,user_id=? where id =?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId, e.ID)
	return err
}
