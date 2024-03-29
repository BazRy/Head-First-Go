package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Parses a float from a string
func GetFloat() (float64, error) {
	reader := bufio.NewReader((os.Stdin))

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	numGrade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return numGrade, nil

}
