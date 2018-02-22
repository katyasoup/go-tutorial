package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// append does not modify; returns new slice that is then assigned to the cards variable
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
