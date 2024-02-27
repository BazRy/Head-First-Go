package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func openFile(fileName string) (*os.File, error) {
	fmt.Printf("Opening file %v\n", fileName)
	return os.Open(fileName)
}
func closeFile(file *os.File) {
	fmt.Printf("Closing file %v\n", file.Name())
	file.Close()
}

func GetFloats(filename string) ([]float64, error) {
	// mydir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mydir)

	file, err := openFile(filename)
	defer closeFile(file)
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

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil

}
