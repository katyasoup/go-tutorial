package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"badass": "#b4d455",
		"white":  "#ffffff",
	}

	printMap(colors)
}

func printMap(myMap map[string]string) {
	for color, hex := range myMap {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// see lecture 51 for differences btwn maps and structs