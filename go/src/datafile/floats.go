package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	// mydir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mydir)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var numbers []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil

}
