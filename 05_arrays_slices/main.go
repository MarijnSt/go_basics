package main

import "fmt"

func main() {
	// arrays
	var fruitArray [2]string

	// waardes toekennen
	fruitArray[0] = "appeltje"
	fruitArray[1] = "peren"

	// declareren en toekennen samen
	fruitsArray := [2]string{"appels", "peren"}

	fmt.Println(fruitsArray)
	fmt.Println(fruitArray[1])

	// slice
	colorsSlice := []string{"green", "blue"}

	fmt.Println(colorsSlice)
	fmt.Println(len(colorsSlice))
}
