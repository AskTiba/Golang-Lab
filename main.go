package main

import "fmt"

// Main function to run the game.
func main() {
	// var bodyType [4]string
	// bodyType[0] = "Slim"
	// bodyType[1] = "Slender"
	// bodyType[2] = "Curvy"

		// var bodyType [4]string

	// bodyType := [4]string{"Slender","Curvy","Petite"}
	// fmt.Println(bodyType)

	var bodyTypes [3][2]string
	bodyTypes[0] = [2]string{"4 petites available", "3 petites booked"}
	bodyTypes[1] = [2]string{"6 slenders available", "1 slender booked"}
	bodyTypes[2] = [2]string{"2 curvys available", "3 curvys booked"}
	fmt.Println("Body type status")

	for i := 0; i < len(bodyTypes); i++ {
		row := bodyTypes[i]
		for j := 0; j < len(row); j++ {
			fmt.Printf("%v ",row[j])
		}
		fmt.Println()
	}
	
}
