package main

import "fmt"

// fibonacci calculates the nth Fibonacci number iteratively.
// The function starts with the first two numbers in the sequence (0 and 1),
// and then iteratively calculates the next numbers up to n.
func fibonacci(n int) int {
	// Base cases: Return the number itself for n = 0 or n = 1.
	if n <= 1 {
		return n
	}

	// Variables to store the two most recent Fibonacci numbers.
	a, b := 0, 1

	// Iteratively calculate the Fibonacci numbers up to the nth number.
	for i := 2; i <= n; i++ {
		// Update 'a' to the current 'b', and 'b' to the sum of 'a' and 'b'.
		a, b = b, a+b
	}

	// Return the nth Fibonacci number.
	return b
}

func main() {
	fmt.Println("Fibonacci sequence:")

	// Print the first 10 Fibonacci numbers (from F(0) to F(9)).
	for i := 0; i < 10; i++ {
		// Call the fibonacci function for each index 'i' and print the result.
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}

	// for loop used as a while loop
	m := 0
	for {
		println(m+1)
		m++

		if m >= 9 {
			break
		}
	}
}
