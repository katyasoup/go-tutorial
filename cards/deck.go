package main

import "fmt"

// Create a new type of "deck" which is a slice of strings (kind of like extension in OOP)
type deck []string

// can now use deck in place of []string!

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
