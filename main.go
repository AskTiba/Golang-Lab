package main

import "fmt"

func sayHello (s string){
	s = "Taa Abwooli"
}

func sayHelloPointer(s *string){

	// Declaring a variable named myStrPointer and it's type is a pointer to a string
	*s = "Taa Abwooli"
}

func main() {
	fmt.Println("What is your favourite programming language?")
	var lang string
	fmt.Scanln(&lang)
	fmt.Printf("Setup %v in your code editor \n", lang)

	greeting := "Taa Akiiki"
	sayHello(greeting)
	fmt.Println("After calling sayHello", greeting)

	sayHelloPointer(&greeting)
	fmt.Println("After calling sayHelloPointer:", greeting)

	// Zero valure of a pointer
	var strPtr *string
	var intPtr *int
	var floatPtr *float64
	var boolPtr *bool 

	fmt.Println(strPtr)
	fmt.Println(intPtr)
	fmt.Println(floatPtr)
	fmt.Println(boolPtr)

}


