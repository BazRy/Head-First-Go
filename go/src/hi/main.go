// This is the main package
package main

import (
	"Head-First-Go/go/src/functionsandpackages"
	"Head-First-Go/go/src/games"
	"Head-First-Go/go/src/keyboard"
	"Head-First-Go/go/src/paint"
	"fmt"
	"log"
)

const PASS_GRADE = 60
const CELCIUS_VALUE = 32
const CELCIUS_MULTIPLIER = 5
const CELCIUS_DIVISOR = 9

// main function
func main() {
	//runFunctionsandpackages()
	//runGames()
	//runCalculatePaintRequired()
	//runPassFail()
	//runCalcTemp()

}

func runFunctionsandpackages() {
	functionsandpackages.Print()
	functionsandpackages.PrintTitleAndFloor()
	functionsandpackages.Reflect()
	functionsandpackages.Apples()
	functionsandpackages.TypeConversion()
	functionsandpackages.GetTheTime()
}
func runPassFail() {
	fmt.Print("Enter a grade: ")
	numGrade, err := keyboard.GetFloat()
	var status string
	if err != nil {
		log.Fatal("*************ERROR OCCURRED*************\n", err)
	}
	if numGrade >= PASS_GRADE {
		status = "Passing"
	} else {
		status = "Failing"
	}

	fmt.Println("A grade of", numGrade, "is a", status, "grade")
}

func runCalcTemp() {
	fmt.Print("Enter a temperature in Farenheit: ")
	farenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal("*************ERROR OCCURRED*************\n", err)
	}
	celcius := (farenheit - CELCIUS_VALUE) * CELCIUS_MULTIPLIER / CELCIUS_DIVISOR

	fmt.Printf("%0.2f degrees celcius\n", celcius)
}

func runGames() {
	games.DoGame()
}

func runCalculatePaintRequired() {

	area := 0.0
	calculatePaintRequired(1.2, 1.8, &area)
	fmt.Printf("Paint required is %.2f litres\n", area)
	area = 0
	calculatePaintRequired(10.5, 11.7, &area)
	fmt.Printf("Paint required is %.2f litres\n", area)
	area = 0
	calculatePaintRequired(10.5, 21.7, &area)
	fmt.Printf("Paint required is %.2f litres\n", area)
}

func calculatePaintRequired(width float64, height float64, area *float64) {

	areaResult, err := paint.PaintNeeded(width, height)

	if err != nil {
		log.Fatal(err)
	}

	*area = areaResult
}
