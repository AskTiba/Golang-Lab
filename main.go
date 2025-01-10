package main

import (
	"fmt"
	"strings"
)

func main() {
	var value float64
	var fromUnit, toUnit string

	// Input: Value to convert
	fmt.Println("Enter value to convert: ")
	fmt.Scanln(&value)

	// Input: From currency
	fmt.Println("Choose unit currency to convert from (Ush, Ksh, USD): ")
	fmt.Scanln(&fromUnit)
	fromUnit = strings.ToUpper(strings.TrimSpace(fromUnit)) // Normalize input

	// Input: To currency
	fmt.Println("Choose unit currency to convert to (Ush, Ksh, USD): ")
	fmt.Scanln(&toUnit)
	toUnit = strings.ToUpper(strings.TrimSpace(toUnit)) // Normalize input

	// Conversion rates between currencies
	exchangeRates := map[string]map[string]float64{
		"USH": {"KSH": 0.038, "USD": 0.00027, "USH": 1},
		"KSH": {"USH": 26.32, "USD": 0.0075, "KSH": 1},
		"USD": {"USH": 3750, "KSH": 133.33, "USD": 1},
	}

	// Validate the currencies
	if rates, ok := exchangeRates[fromUnit]; ok {
		if rate, ok := rates[toUnit]; ok {
			// Perform conversion
			convertedValue := value * rate
			fmt.Printf("%.2f %s is equal to %.2f %s\n", value, fromUnit, convertedValue, toUnit)
			return
		}
	}

	// Handle invalid input
	fmt.Println("Invalid currency conversion. Please choose Ush, Ksh, or USD.")
}
