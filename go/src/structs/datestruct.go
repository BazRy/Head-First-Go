package main

import "fmt"

const MIN int = 1
const MAX_MONTH int = 12
const MAX_DAY int = 31

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{}
	date.SetYear(2024)
	date.SetMonth(2)
	date.SetDay(23)
	fmt.Println(date)
}

func (d *Date) SetYear(year int) {
	d.Year = year
}
func (d *Date) SetMonth(month int) {
	d.Month = month
}
func (d *Date) SetDay(day int) {
	d.Day = day
}
