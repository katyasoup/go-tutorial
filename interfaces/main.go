package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// this is dumb and redundant, but accomplishes what we actually want to do:
func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for generating English greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// custom logic for generating Spanish greeting
	return "Hola!"
}

// can omit eb/sb value since not used in function
