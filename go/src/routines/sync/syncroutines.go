package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("recieving go routine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}

func reportNap(name string, delay int) {

	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}

	fmt.Println(name, "wakes up")
}

func send(myChannel chan string) {
	reportNap("sending go routine", 2)
	fmt.Println("***sending value a***")
	myChannel <- "a"
	fmt.Println("***sending value b***")
	myChannel <- "b"
}
