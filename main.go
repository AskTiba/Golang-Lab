package main

import "fmt"

// Main function to run the game.
func main() {
	colors := [4]string{"Green", "Blue","Orange", "Purple"}
	fmt.Println(colors)
	for i, color := range colors{
		fmt.Printf("index:%v. Item:%v\n", i, color)
		fmt.Println()
	}

	// If you only want the item, then you can discard thr index
	for i := range colors{
		fmt.Printf("%v\n",i)
		fmt.Println()
	}
	
	// If you only want the index, then you can discard thr item
	for _, color := range colors{
		fmt.Printf("%s\n",color)
		fmt.Println()
	}

	gotChars := map[string]string{
		"Tomund" : "Wildling",
		"Brienne" : "Tarth",
		"Sam" : "Tally",
		"Aegon" : "Targaryan",
		"Robert" : "Baratheon",
		"Aary" : "Stark",
	}

	for name,house := range gotChars{
		fmt.Printf("%s of house %s\n", name, house)
		fmt.Println()
	}

	for _, house := range gotChars{
		fmt.Printf("%s\n", house)
	}

	// // Total number gotChars
	// totalNumberOfGotChars := 0
	// for _, house := range gotChars{

	// }
}