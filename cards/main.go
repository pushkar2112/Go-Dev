package main

import "fmt"

func main() {

	cards := []string{newCard(), "Ace of Diamonds"}

	cards = append(cards, "Six of Spades") // append returns a new slice of strings
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string { // function returns a string
	return "Five of Diamonds"
}