package main

import "fmt"

func main() {

	card := newCard()

	fmt.Println(card)
}

func newCard() string { // function returns a string
	return "Five of Diamonds"
}