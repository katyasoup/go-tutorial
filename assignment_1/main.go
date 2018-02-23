package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if number%2 != 0 {
			// can also do for i... if numbers[i]
			fmt.Println(number, "is odd")
		} else {
			fmt.Println(number, "is even")
		}
	}

}
