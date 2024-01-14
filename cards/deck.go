package main

import (
	"fmt"
	"os"
	"strings" // string manipulation
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string // deck is a slice of strings

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // slice of strings
	cardValues := []string{"Ace", "Two", "Three", "Four"} // slice of strings

	for _, suit := range cardSuits { // iterate through cardSuits
		for _, value := range cardValues { // iterate through cardValues
			cards = append(cards, value + " of " + suit) // append returns a new slice of strings
		}
	}

	return cards
}

func (d deck) print() { // receiver function
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // convert deck to slice of strings then to string and join with comma separator
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is the permission to write to file
	// ioutil is deprecated, os provides the same functionality
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) // returns byte slice and error
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // exit program
	}

	s := strings.Split(string(bs), ",") // convert byte slice to string then to slice of strings

	return deck(s) // convert slice of strings to deck
}