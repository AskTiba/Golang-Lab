package main

import "fmt"

func main() {
	const myname string = "Tibamwenda Anthony"
	fmt.Println(myname)

	// With multiple constants
	const (
		place_of_birth = "Rubaga"
		founded = 2008
		founder = "Karamagi Nelson"
	)

	fmt.Println("Founder:", founder)
	
	const (
		_ = iota
		Currency
		Language
		Country
		City
		President
	)
	fmt.Println("Currency", Currency)
	fmt.Println("Language", Language)
	fmt.Println("Country", Country)
	fmt.Println("City", City)
	fmt.Println("President", President)
}


