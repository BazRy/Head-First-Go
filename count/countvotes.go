package main

import (
	"Head-First-Go/datafile"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	//candidates, err := datafile.GetStrings("votes.txt")
	candidates, err := datafile.GetStrings(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	votesPerCandidate := make(map[string]int)

	for _, candidate := range candidates {
		votesPerCandidate[candidate]++
	}

	var names []string
	for key := range votesPerCandidate {
		names = append(names, key)
	}
	sort.Strings(names)

	fmt.Printf("Votes are as follows:\n")
	for index, name := range names {
		plural := ""
		value := votesPerCandidate[name]
		if value > 1 {
			plural = "s"
		}
		fmt.Printf("%v) %v has %v vote%v\n", index+1, name, value, plural)
	}

}
