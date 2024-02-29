package main

import (
	"fmt"
	"log"
	"net/http"
)

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, []byte("Hello from the internet"))
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, []byte("Salut from the internet"))

}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, []byte("Namaste from the internet"))

}

func write(writer http.ResponseWriter, message []byte) {
	_, err := writer.Write(message)

	if err != nil {
		log.Fatal(err)
	}
}

func sayHi() {
	fmt.Println("Hi")
}
func greet(myFunc func()) {
	myFunc()
}

func printSum(num1 int, num2 int) bool {
	fmt.Println(num1 * num2)
	return true
}

func main() {
	//example first class function assignment and usage
	myFunc := sayHi
	greet(myFunc)

	//long handed way to declare the function
	//var myFunc2 func(int, int) bool
	myFunc2 := printSum
	resp := myFunc2(2, 3)
	fmt.Println("Response is", resp)

	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
