package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// // see deal function in deck.go, returns two values of type deck, so assign to two vars

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())
}
