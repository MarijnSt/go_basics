package main

import "fmt"

func main() {
	// met var keyword
	var name = "Marijn"
	var age = 27
	var isCool = true

	// aanpassen
	isCool = false

	// shorthand
	score := 3.5

	// meerdere samen
	adres, mail := "Bambrugge", "marijn@gmail.com"

	fmt.Println(name, age, isCool, score, adres, mail)
	fmt.Printf("%T\n", score)
}
