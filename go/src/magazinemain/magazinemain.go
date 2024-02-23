package main

import (
	"Head-First-Go/go/src/magazinemain/subscription"
)

// get a pointer to a struct, not a copy.
// Update the struct using dot notation
// pass the pointer to be printed.
func main() {

	sub1 := subscription.GetDefaultSubscriber("Barry Ryan")
	subscription.PrintInfo(sub1)

	//struct literal
	sub2 := subscription.Subscriber{Name: "Millie Ryan", Rate: 123, Active: true}
	subscription.PrintInfo(&sub2) //in this case subscriber is not a pointer so have to pass arg as a poiner
}
