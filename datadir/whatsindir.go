package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	defer recoverFromPanic()
	dir, _ := os.Getwd()
	scanDirectory(dir)

}

func scanDirectory(path string) {

	fmt.Printf("Starting to scan %v\n", path)

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		// if file.Name() == "subdir3" {
		// 	panic("something unfortunate")
		// }

		isDir := file.IsDir()
		if isDir {
			filePath := filepath.Join(path, file.Name())
			scanDirectory(filePath)
		} else {
			fmt.Printf("Found file named %v\n", file.Name())
		}
	}
}

func recoverFromPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(string)
	if ok {
		if err == "disaster" {
			fmt.Printf("A disaster has occurred\n")
		} else {
			panic("Reinstating panic " + err)
		}
	}
}
