package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

// must tell Go what type of data the function will return - so now Go knows that the card variable will always hold a string
