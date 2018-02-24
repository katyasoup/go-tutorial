package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://katyasoup.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Hmm, looks like", link, "might be down :(")
		c <- "Channel says: Oof, might be down"
		return
	}

	fmt.Println("All gravy,", link, "is up! :)")
	c <- "Channel says: Yep, it's up"
}
