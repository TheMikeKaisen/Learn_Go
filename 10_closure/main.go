package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++ // Accesses and modifies the outer variable `count`
		return count
	}
}

func main() {
	increment := counter() // `increment` is a closure

	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3

	// Create a new counter
	newCounter := counter()
	fmt.Println(newCounter()) // Output: 1
}
