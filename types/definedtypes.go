package main

import (
	"fmt"
)

type Litres float64
type Gallons float64
type Pints float64

const LITRES_TO_GALLONS_MULTIPLIER float64 = 0.264
const GALLONS_TO_LITRES_MULTIPLIER float64 = 3.785
const PINTS_TO_LITRES_MULTIPLIER float64 = 0.568

func main() {

	gallons := Gallons(10)
	litres := Litres(100)
	pints := Pints(10)

	fmt.Printf("Litres: %0.2f, Gallons: %0.2f\n", litres, gallons)
	fmt.Printf("Add together: %0.2f\n", litres+Litres(gallons))
	fmt.Printf("Litres to Gallons: %0.2f Litres = %0.2f Gallons\n", litres, litres.toGallons())
	fmt.Printf("Gallons to Litres: %0.2f Gallons = %0.2f Litres\n", gallons, gallons.toLitres())
	fmt.Printf("Pints to Litres: %0.2f Pints = %0.2f Litres\n", pints, pints.toLitres())
	pints.double()
	fmt.Printf("Pints doubled: %0.2f Pints\n", pints)
}

// we cannot define functions with the same name (even overloaded like in java)
// but we can create methods with the same name (functions with a new reciever param)
func (l Litres) toGallons() Gallons {
	return Gallons(l * Litres(LITRES_TO_GALLONS_MULTIPLIER))
}

func (g Gallons) toLitres() Litres {
	return Litres(g * Gallons(GALLONS_TO_LITRES_MULTIPLIER))
}

func (p Pints) toLitres() Litres {
	return Litres(p * Pints(PINTS_TO_LITRES_MULTIPLIER))
}

func (p *Pints) double() {
	*p = *p * Pints(2)
}
