package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var numbers []float64

	for _, arg := range arguments {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

func average(numbers ...float64) float64 {

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}
