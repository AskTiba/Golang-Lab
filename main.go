package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	pi := math.Pi
	phi := math.Phi

	fmt.Println("The value of Pi is:", pi)
	fmt.Println("The value of Phi is:", phi)

	// Integer to float
	distanceInt := 42
	var distanceFloat float64 = float64(distanceInt)
	fmt.Printf("Integer to float: %.3f\n", distanceFloat)

	// Checking the type of a variable
	fmt.Printf("%T\n", phi)

	// Convert a variable to a string
	str := fmt.Sprint(phi)
	fmt.Println("The string value of Phi is:", str)
	fmt.Printf("%T\n", str)

	// int to string
	str1 := strconv.Itoa(25)
	fmt.Println(str1)
	fmt.Printf("%T\n", str1)

	// sting to in
	str2 := "49"
	var intFromStr, _ = strconv.Atoi(str2)
	fmt.Println("String to integer:", intFromStr)
}
