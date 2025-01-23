package main

import "fmt"

// Main function to run the game.
func main() {
	// carInventory := make(map[string]int)
	// carInventory["Nissan"] = 6
	// carInventory["Isuzu"] = 3
	// carInventory["Subaru"] = 13
	// carInventory["Jeep"] = 6
	// carInventory["Benz"] = 6

	carInventory := map[string]int{
		"Nissan" : 6,
		"Isuzu" : 3,
		"Subaru": 13,
		"Jeep" : 6,
		"Benz" : 8,
	}

	fmt.Printf("Car inventory is: %v\n", carInventory)
	cars := len(carInventory)
	fmt.Println()

	// lens returns the number of key-vakue pairs
	fmt.Printf("There are %v types of cars\n",cars)
	fmt.Println()
	
	//  update the map
	carInventory["Benz"] = 15
	fmt.Printf("Car inventory is: %v\n", carInventory)
	fmt.Println()
	
	// add to the map
	carInventory["Ferrari"] = 7
	carInventory["Ferrari"] = 7
	fmt.Printf("Car inventory is: %v\n", carInventory)
	fmt.Println()
	
	// Delete item from the map the map
	delete(carInventory,"Isuzu")
	fmt.Printf("Car inventory is: %v\n", carInventory)
	fmt.Println()

	// retrieving a vakue from a map using the ghost key syntax
	numFerraris, FerrariFound := carInventory["Ferrari"]
	fmt.Printf("Ferraris found: %t\n", FerrariFound)
	if FerrariFound{
		fmt.Printf("We have %v Ferraris\n",numFerraris)
	}
	fmt.Println()
	
	//  OR
	
	// This is a common practice in Go and it is useful  when the variables are only
	// needed  for the check, they are only scoped here
	if numFerraris, ok := carInventory["Ferrari"]; ok{fmt.Printf("We have %v Ferraris\n",numFerraris)}
	fmt.Println()
	
	clear(carInventory)
	fmt.Printf("Car inventory is: %v\n", carInventory)
}
