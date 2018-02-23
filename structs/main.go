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
	contact   contactInfo
}

func main() {
	// katie := person{firstName: "Katie", lastName: "Campbell", age: 29}
	// fmt.Println(katie)

	var katie person
	// declare a struct with no assigned values

	katie.firstName = "Katie"
	katie.lastName = "Campbell"
	katie.age = 29
	// create attributes for katie struct

	fmt.Println(katie)
	fmt.Printf("%+v", katie)

}
