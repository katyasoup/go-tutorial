package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	// cards.shuffle()
	cards.print()
	fmt.Println(len(cards))
}
