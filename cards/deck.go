package main

import "fmt"

// Create a new type of "deck" which is a slice of strings (kind of like extension in OOP)
type deck []string

// can now use deck in place of []string!

// should create and return a list of all the cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	// replace i, j, etc with _ when we need to declare the var and not use it

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	// any variable of type "deck" now has access to print method
	// d in this case is kind of similar to this/self
	// usually use 1- or 2-letter abbreviation related to the type - but can call it whatever you want
}
