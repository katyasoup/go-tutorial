package main

import "fmt"

// Create a new type of "deck" which is a slice of strings (kind of like extension in OOP)
type deck []string

// can now use deck in place of []string!

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	// any variable of type "deck" now has access to print method
	// d in this case is kind of similar to this/self
	// usually use 1- or 2-letter abbreviation related to the type - but can call it whatever you want
}
