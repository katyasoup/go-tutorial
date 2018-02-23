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
	katie := person{
		firstName: "Katie",
		lastName:  "Campbell",
		age:       29,
		contact: contactInfo{
			email:   "hello@katyasoup.com",
			zipCode: 55405,
		}, // needs a comma too!
	}
	fmt.Printf("%+v", katie)
}
