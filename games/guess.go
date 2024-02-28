package games

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func DoGame() {

	winner := false
	numAttempts := 10

	target := rand.Intn(100) + 1
	fmt.Println("I've selected a number beetween 1 & 100")
	fmt.Println("Can you guess it?")

	for numAttempts > 0 {
		fmt.Printf("You have %v attempts remaining\n", numAttempts)
		fmt.Println("Enter a number: ")
		reader := bufio.NewReader((os.Stdin))

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if num == target {
			fmt.Printf("Excellent you were spot on!\n")
			winner = true
			break
		} else if num > target {
			fmt.Printf("Sorry, you were too high\n")
		} else {
			fmt.Printf("Sorry, you were too low\n")
		}
		numAttempts--
	}

	if !winner {
		fmt.Printf("Sorry, you didn't win on this occasion, number was %v\n", target)
	}

}
