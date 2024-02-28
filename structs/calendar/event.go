package calendar

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type Event struct {
	name string
	date Date
}

func SetUpEvent(eventName string, year int, month int, day int) (Event, error) {
	event := Event{}
	eventDate := Date{}

	err := event.setName(eventName)
	if err == nil {
		err = eventDate.setYear(year)
	}
	if err == nil {
		err = eventDate.setMonth(month)
	}
	if err == nil {
		err = eventDate.setDay(day)
	}
	if err == nil {
		event.date = eventDate
	}
	return event, err
}

func (e *Event) setName(eventName string) error {
	eventNameLen := utf8.RuneCountInString(strings.TrimSpace(eventName))
	if eventNameLen > 15 || eventNameLen == 0 {
		return errors.New("invalid event name")
	}
	e.name = eventName
	return nil
}

func (e Event) Name() string {
	return e.name
}
func (e Event) Month() int {
	return e.date.month
}
func (e Event) Year() int {
	return e.date.year
}
func (e Event) Day() int {
	return e.date.day
}
