package main

import (
	"Head-First-Go/structs/calendar"
	"fmt"
	"log"
)

func main() {
	event, err := calendar.SetUpEvent("Barry's Event", 2024, 12, 28)
	failOnError(err, "An error has occurred")
	fmt.Printf("Event details:\n Name: %v\n Date: %v/%v/%v\n", event.Name(), event.Day(), event.Month(), event.Year())
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
