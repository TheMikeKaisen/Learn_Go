package main

import "fmt"

func main() {
	x := 42
	ptr := &x // Get the memory address of x

	fmt.Println("Value of x:", x)       // Output: 42
	fmt.Println("Address of x:", ptr)   // Output: Address (e.g., 0xc0000140b0)
	fmt.Println("Value via ptr:", *ptr) // Output: 42 (dereferencing)

	// Modify the value via the pointer
	*ptr = 50
	fmt.Println("Updated value of x:", x) // Output: 50
}
