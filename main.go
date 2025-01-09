package main

import (
	"fmt"
)

func main() {
	// Variables representing conditions
	isSunny := false
	isFree := false

	// Print the state of conditions
	fmt.Println("Is it sunny? ", isSunny)
	fmt.Println("Are you free? ", isFree)

	// Logical conditions using AND (&&)
	if isSunny && isFree {
		fmt.Println("It's sunny and you're free! Let's go to the park.")
	}

	// Logical conditions using OR (||)
	if isSunny || isFree {
		fmt.Println("Either it's sunny or you're free (or both). Consider going to the park.")
	}

	// Logical conditions using NOT (!)
	if !isSunny {
		fmt.Println("It's not sunny today.")
	}
	if !isFree {
		fmt.Println("You are not free today.")
	}

	// Combining NOT with AND
	if !isSunny && !isFree {
		fmt.Println("It's neither sunny nor are you free. No park today.")
	}

	// Combining NOT with OR
	if !(isSunny || isFree) {
		fmt.Println("Neither sunny nor free (checked with OR and NOT). No park today.")
	}

	// Complex condition: Free, but weather is bad
	if isFree && !isSunny {
		fmt.Println("You're free, but it's not sunny. Maybe do something indoors.")
	}

	// Complex condition: Sunny, but you're busy
	if isSunny && !isFree {
		fmt.Println("It's sunny, but you're busy. Maybe another time.")
	}
}
