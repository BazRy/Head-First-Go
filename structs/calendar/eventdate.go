package calendar

import (
	"errors"
)

const MIN int = 1
const MAX_MONTH int = 12
const MAX_DAY int = 31

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) setYear(year int) error {
	if year < MIN {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) setMonth(month int) error {
	if month < MIN || month > MAX_MONTH {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) setDay(day int) error {

	if day < MIN || day > MAX_DAY {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
