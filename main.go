package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("What is your current speed?")
	var speed string
	fmt.Scanln(&speed)
	fmt.Println("Your speed is:", speed, "Km/hr")
	fmt.Printf("%T\n", speed)
	var speedInt, _ = strconv.Atoi(speed)
	fmt.Printf("%T\n", speedInt)

	if (speedInt>=130) {
		fmt.Println("Slow down, you're driving too fast!!!")
	}else if (speedInt>=45) {
		fmt.Println("Welcome to Fort Potal, drive safe.")
	}else{
		fmt.Println("Your speed is dangerously on this highway,speed up")
	}
	
}


