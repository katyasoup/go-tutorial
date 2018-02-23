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

	katiePointer := &katie
	katiePointer.updateName("Zeke")
	katie.print()
	// see section 4, lecture 43/44 for explanation of pointers
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
