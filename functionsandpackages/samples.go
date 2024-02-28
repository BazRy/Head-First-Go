package functionsandpackages

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
)

func Print() {
	fmt.Println("Functions and packages sample")
}

func PrintTitleAndFloor() {
	fmt.Println(math.Floor(1.03))
	fmt.Println(strings.Title("a busy bee"))
}

func Reflect() {
	fmt.Println(reflect.TypeOf("xsd"))
	fmt.Println(reflect.TypeOf(-1))
	fmt.Println(reflect.TypeOf(1.25))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf('A'))
}

func Variables() {
	var firstName, lastName string = "Barry", "Ryan"
	fmt.Println(firstName + lastName)
}

func Apples() {
	originalCount := 10
	eatenCount := 4

	fmt.Printf("I started with %v %v\n", originalCount, "apples")
	fmt.Printf("Some jerk ate %v %v\n", eatenCount, "apples")
	fmt.Printf("Now there are %v %v\n", originalCount-eatenCount, "apples")

}

func TypeConversion() {
	age := 20
	day := 1.5

	fmt.Println(float64(age) * day)
	fmt.Println(age * int(day))
}

func GetTheTime() {
	now := time.Now()
	fmt.Println(now.Year())
}
