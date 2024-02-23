package subscription

import (
	"fmt"
)

func PrintSubcriberInfo() {
}

// return a pointer not a copy
func GetDefaultSubscriber(name string) *Subscriber {
	var sub1 Subscriber
	sub1.Name = name
	sub1.Rate = 725
	sub1.Active = true
	sub1.Address = Address{Firstline: "400 Thorpe Road", City: "Peterborough", Postcode: "PE3 6NB"}
	return &sub1
}

// pass pointers for large structs as otherwise we'll be holding the original struct and a copy in memory
func PrintInfo(sub *Subscriber) { //accept a pointer as an argument
	fmt.Println("Name: ", sub.Name)
	fmt.Println("Rate: ", sub.Rate)
	fmt.Println("Active: ", sub.Active)
	fmt.Printf("Address: %v, %v, %v\n", sub.Firstline, sub.City, sub.Postcode) //anonymous struct can have it's fields access as if part of outer struct
}

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address //struct within a struct
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address //inner struct,  it's fields have to be access by: <name of stuct var>.HomeAddress.City = "ggv"
}

type Address struct {
	Firstline string
	City      string
	Postcode  string
}
