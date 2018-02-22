package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// append does not modify; returns new slice that is then assigned to the cards variable

	for i, card := range cards {
		fmt.Println(i, card)
	}
	// "range" is the keyword to iterate over every record inside of a slice
	// use := to reinitialize variable with each iteration
}

func newCard() string {
	return "Five of Diamonds"
}
