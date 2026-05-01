package models

import (
	"log/slog"
	"time"
)

type Event struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"  binding:"required,min=3"`
	Description string    `json:"description"  binding:"required,min=3"`
	Location    string    `json:"location"  binding:"required"`
	DateTime    time.Time `json:"datetime"  binding:"required"`
	UserId      uint64    `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() {
	/* Here WE Will Add the Db Code */
	events = append(events, e)
	slog.Info("new Event Added here")
}

func GetAllEvents() []Event {
	slog.Info("fetcging all the events", "events", events)
	return events
}
