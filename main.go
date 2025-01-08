package main

import (
	"fmt"
	"math"
)

func main() {
	phi := math.Phi
	fmt.Println("The value of phi is:",phi)
	fmt.Println("phi (Round):", math.Round(phi))
	fmt.Println("phi (Ceil):", math.Ceil(phi))
	fmt.Println("phi (Floor):", math.Floor(phi))

	besty := math.Abs(-13) 
	fmt.Println("The absolute value of -13 is:", besty)

	exponent := math.Pow(2,16)
	fmt.Println("2 exponent 16 is:", exponent)

	squareRoot := math.Sqrt(144)
	fmt.Println("The square root of 144 is:",squareRoot)
}
