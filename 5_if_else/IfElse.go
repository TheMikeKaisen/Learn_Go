package main

import "fmt"

func main() {
	num := 15

	// Basic if-else
	if num%2 == 0 {
		fmt.Println("Even")
	} else if num%3 == 0 {
		fmt.Println("Divisible by 3")
	} else {
		fmt.Println("Odd and not divisible by 3")
	}

	// If with initialization
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is 5 or less")
	}
}
