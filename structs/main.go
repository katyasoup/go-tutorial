package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	katie := person{
		firstName: "Katie",
		lastName:  "Campbell",
		age:       29,
		contactInfo: contactInfo{
			email:   "hello@katyasoup.com",
			zipCode: 55405,
		},
	}

	katie.updateName("Zeke")
	katie.print()
	// go allows this shortcut to automatically turn person into pointerToPerson, neat!
	// but you GOTTA keep that asterisk in the receiver down there
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// note - slices are modified on the original, unlike structs! GOTCHA
// (explanation in lecture 47)
