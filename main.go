package main

import (
	"fmt"
	"strings"
)

func main() {
    house := "Stark"
    name := "Red"
	words := "Winter is coming!"
    fmt.Println("Hello",name, house, ".Tell them that", words, "!")

	IsPresent := true
	fmt.Println("It is", !IsPresent, "that United lost today.")

	str1 := "Welcome to Jamaica"
	str2 := "welcome to jamaica"
	fmt.Println(strings.EqualFold(str1,str2))

	fmt.Println(strings.Index(name,`d`))
	fmt.Println(len(house))

	str3 := "Dorothy"
	isIndex := strings.Index(words,`is`)
	fmt.Println(words[isIndex+3:16])
	// fmt.Println(strings.Index(str3,`y`))
	fmt.Println(str3[:4])
	
	// Numbers, F = 9C/5 + 32
	// F = (K-273.15)*1.8+32, where K = 288.15
	temperatureK := 288.15
	temperatureF := 1.8*(temperatureK-273.15) + 32
	fmt.Println("Temperature in Kelvin:", temperatureK)
	fmt.Println("Temperature in Fahreheit:", temperatureF)

	// when F = 70, K = ?
	temperatureFah := 70.0
	temperatureKel := ((temperatureFah-32)/(1.8))+273.15
	fmt.Println("Temperature in Fahreheit:", temperatureFah)
	fmt.Println("Temperature in Kelvin:", temperatureKel)


}
