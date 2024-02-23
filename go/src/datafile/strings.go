package datafile

import (
	"bufio"
	"log"
	"os"
)

func GetStrings(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return lines, err
		}
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil

}
