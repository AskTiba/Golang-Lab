package main

import "fmt"

// Main function to run the game.
func main() {
	fuelTypes := []string{"Gas","Electric","Hybrid","Others"}
	fuelTypes = append(fuelTypes, "Diesel")
	fmt.Println(len(fuelTypes))
	fmt.Println(fuelTypes)
	fmt.Println("The slice is nill:", fuelTypes == nil)
	fmt.Println()

	iceCreamFlavour := make([]string,2)
	fmt.Println(len(iceCreamFlavour))
	fmt.Println("The slice is nill:", iceCreamFlavour == nil)
	fmt.Println(iceCreamFlavour)
	iceCreamFlavour[0] = "Strawberry"
	iceCreamFlavour[1] = "Pistachio"
	iceCreamFlavour = append(iceCreamFlavour,"Vanilla","Chocholate","Hazelnut")
	fmt.Println(len(iceCreamFlavour))
	fmt.Println(iceCreamFlavour)

}
