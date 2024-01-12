package main

import "fmt"

func main() {

	cards := []string{newCard(), "Ace of Diamonds"}

	cards = append(cards, "Six of Spades") // append returns a new slice of strings
	
	fmt.Println(cards)	// slice of strings
}

func newCard() string { // function returns a string
	return "Five of Diamonds"
}