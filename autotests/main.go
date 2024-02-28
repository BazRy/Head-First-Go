package main

import (
	"Head-First-Go/autotests/prose"
	"fmt"
)

func main() {
	phrases := []string{"Pearl", "Bella", "Millie"}
	fmt.Println(prose.JoinWithCommas(phrases))
	phrases = []string{"Barry", "Janey"}
	fmt.Println(prose.JoinWithCommas(phrases))

}
