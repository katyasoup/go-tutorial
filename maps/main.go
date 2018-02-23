package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string

	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	delete(colors, "white")

	// colors := map[string]string{
	// 	"red":       "#ff0000",
	// 	"coolGreen": "#b4d455",
	// }

	fmt.Println(colors)
}
